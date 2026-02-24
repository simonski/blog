package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	rendererhtml "github.com/yuin/goldmark/renderer/html"
)

const (
	postsDir      = "posts"
	draftsDir     = "drafts"
	outputDir     = "output"
	templatesDir  = "templates"
	indexCSSName  = "index.css"
	postCSSName   = "post.css"
	indexFileName = "index.html"
)

var (
	mdLinkPattern = regexp.MustCompile(`!?\[[^\]]*\]\(([^)]+)\)`)
	datePattern   = regexp.MustCompile(`\b\d{4}-\d{2}-\d{2}\b`)
	mdEngine      = goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(rendererhtml.WithUnsafe()),
	)
)

type templates struct {
	PostHeader  string
	PostFooter  string
	IndexHeader string
	IndexFooter string
}

type post struct {
	Title      string
	DateRaw    string
	Date       time.Time
	Body       string
	Slug       string
	SourcePath string
	IsDraft    bool
}

func main() {
	includeDrafts := flag.Bool("draft", false, "include posts from drafts")
	flag.Parse()

	if err := buildSite(*includeDrafts); err != nil {
		fmt.Fprintf(os.Stderr, "build failed: %v\n", err)
		os.Exit(1)
	}
}

func buildSite(includeDrafts bool) error {
	tmpl, err := loadTemplates()
	if err != nil {
		return err
	}

	if err := resetOutputDir(); err != nil {
		return err
	}
	if err := copyStylesheets(); err != nil {
		return err
	}

	posts, err := collectPosts(postsDir, false)
	if err != nil {
		return err
	}
	if includeDrafts {
		drafts, draftErr := collectPosts(draftsDir, true)
		if draftErr != nil {
			return draftErr
		}
		posts = append(posts, drafts...)
	}

	assignUniqueSlugs(posts)

	for _, p := range posts {
		if err := renderPost(p, tmpl); err != nil {
			return err
		}
	}

	sortPosts(posts)
	if err := renderIndex(posts, tmpl); err != nil {
		return err
	}

	return nil
}

func loadTemplates() (templates, error) {
	postHeader, err := readTemplate(filepath.Join(templatesDir, "post_header.html"))
	if err != nil {
		return templates{}, err
	}
	postFooter, err := readTemplate(filepath.Join(templatesDir, "post_footer.html"))
	if err != nil {
		return templates{}, err
	}
	indexHeader, err := readTemplate(filepath.Join(templatesDir, "index_header.html"))
	if err != nil {
		return templates{}, err
	}
	indexFooter, err := readTemplate(filepath.Join(templatesDir, "index_footer.html"))
	if err != nil {
		return templates{}, err
	}

	return templates{
		PostHeader:  postHeader,
		PostFooter:  postFooter,
		IndexHeader: indexHeader,
		IndexFooter: indexFooter,
	}, nil
}

func readTemplate(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read template %s: %w", path, err)
	}
	return string(data), nil
}

func resetOutputDir() error {
	if err := os.RemoveAll(outputDir); err != nil {
		return fmt.Errorf("reset output directory: %w", err)
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	return nil
}

func copyStylesheets() error {
	files := []string{indexCSSName, postCSSName}
	for _, name := range files {
		src := filepath.Join(templatesDir, name)
		dst := filepath.Join(outputDir, name)
		if err := copyFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func collectPosts(root string, isDraft bool) ([]post, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("read directory %s: %w", root, err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	var posts []post
	for _, entry := range entries {
		if entry.IsDir() {
			postPath, postErr := findSingleMarkdown(filepath.Join(root, entry.Name()))
			if postErr != nil {
				return nil, postErr
			}
			if postPath == "" {
				continue
			}

			parsed, parseErr := parsePost(postPath, entry.Name(), isDraft)
			if parseErr != nil {
				return nil, parseErr
			}
			posts = append(posts, parsed)
			continue
		}

		if isDraft && strings.HasSuffix(strings.ToLower(entry.Name()), ".md") {
			sourcePath := filepath.Join(root, entry.Name())
			slug := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
			parsed, parseErr := parsePost(sourcePath, slug, true)
			if parseErr != nil {
				return nil, parseErr
			}
			posts = append(posts, parsed)
		}
	}

	return posts, nil
}

func findSingleMarkdown(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("read post directory %s: %w", dir, err)
	}

	var matches []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToLower(entry.Name()), ".md") {
			matches = append(matches, filepath.Join(dir, entry.Name()))
		}
	}

	switch len(matches) {
	case 0:
		return "", nil
	case 1:
		return matches[0], nil
	default:
		sort.Strings(matches)
		return "", fmt.Errorf("expected one markdown file in %s, found %d", dir, len(matches))
	}
}

func parsePost(path string, slug string, isDraft bool) (post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return post{}, fmt.Errorf("read markdown %s: %w", path, err)
	}

	title, dateRaw, body := parseFrontMatter(string(data))
	if title == "" {
		title = inferTitle(body, path)
	}
	if dateRaw == "" {
		dateRaw = inferDate(string(data))
	}

	parsedDate, err := time.Parse("2006-01-02", dateRaw)
	if err != nil {
		parsedDate = time.Time{}
	}

	return post{
		Title:      title,
		DateRaw:    dateRaw,
		Date:       parsedDate,
		Body:       body,
		Slug:       sanitizeSlug(slug),
		SourcePath: path,
		IsDraft:    isDraft,
	}, nil
}

func parseFrontMatter(markdown string) (title string, date string, body string) {
	content := strings.ReplaceAll(markdown, "\r\n", "\n")
	content = strings.TrimPrefix(content, "\uFEFF")
	lines := strings.Split(content, "\n")

	start := 0
	for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	if start >= len(lines) {
		return "", "", ""
	}

	metadata := map[string]string{}
	bodyStart := start

	if strings.TrimSpace(lines[start]) == "---" {
		bodyStart = start + 1
		for i := start + 1; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])
			if line == "" || line == "---" {
				bodyStart = i + 1
				break
			}
			key, value, ok := parseMetadataLine(lines[i])
			if !ok {
				continue
			}
			metadata[key] = value
			bodyStart = i + 1
		}
	} else {
		key, value, ok := parseMetadataLine(lines[start])
		if !ok || (key != "title" && key != "date") {
			return "", "", content
		}
		metadata[key] = value
		bodyStart = start + 1
		for i := start + 1; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])
			if line == "" {
				bodyStart = i + 1
				break
			}
			key, value, ok := parseMetadataLine(lines[i])
			if !ok {
				bodyStart = i
				break
			}
			metadata[key] = value
			bodyStart = i + 1
		}
	}

	bodyLines := append([]string(nil), lines[bodyStart:]...)
	for len(bodyLines) > 0 && isBodySeparatorLine(bodyLines[0]) {
		bodyLines = bodyLines[1:]
	}
	for len(bodyLines) > 0 && strings.TrimSpace(bodyLines[len(bodyLines)-1]) == "" {
		bodyLines = bodyLines[:len(bodyLines)-1]
	}
	if len(bodyLines) > 0 && strings.TrimSpace(bodyLines[len(bodyLines)-1]) == "---" {
		bodyLines = bodyLines[:len(bodyLines)-1]
	}

	return metadata["title"], metadata["date"], strings.Join(bodyLines, "\n")
}

func parseMetadataLine(line string) (key string, value string, ok bool) {
	rawKey, rawValue, found := strings.Cut(line, ":")
	if !found {
		return "", "", false
	}

	key = strings.ToLower(strings.TrimSpace(rawKey))
	value = strings.TrimSpace(rawValue)
	if key == "" {
		return "", "", false
	}

	for _, r := range key {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' && r != '_' {
			return "", "", false
		}
	}

	return key, value, true
}

func isBodySeparatorLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) < 3 {
		return false
	}
	for _, r := range trimmed {
		if r != '-' {
			return false
		}
	}
	return true
}

func inferTitle(markdown string, path string) string {
	lines := strings.Split(strings.ReplaceAll(markdown, "\r\n", "\n"), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
		}
	}
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func inferDate(markdown string) string {
	match := datePattern.FindString(markdown)
	return match
}

func sanitizeSlug(slug string) string {
	slug = strings.TrimSpace(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.Trim(slug, "/")
	if slug == "" {
		return "post"
	}
	return slug
}

func assignUniqueSlugs(posts []post) {
	seen := map[string]int{}
	for i := range posts {
		original := posts[i].Slug
		count := seen[original]
		if count == 0 {
			seen[original] = 1
			continue
		}
		count++
		seen[original] = count
		posts[i].Slug = fmt.Sprintf("%s-%d", original, count)
	}
}

func renderPost(p post, tmpl templates) error {
	var rendered bytes.Buffer
	if err := mdEngine.Convert([]byte(p.Body), &rendered); err != nil {
		return fmt.Errorf("render markdown for %s: %w", p.SourcePath, err)
	}

	postDir := filepath.Join(outputDir, p.Slug)
	if err := os.MkdirAll(postDir, 0o755); err != nil {
		return fmt.Errorf("create post output directory %s: %w", postDir, err)
	}

	if err := copyReferencedImages(p.Body, filepath.Dir(p.SourcePath), postDir); err != nil {
		return err
	}

	relToRoot, err := filepath.Rel(postDir, outputDir)
	if err != nil {
		return fmt.Errorf("resolve css path for %s: %w", p.Slug, err)
	}
	cssHref := postCSSName
	if relToRoot != "." {
		cssHref = filepath.ToSlash(filepath.Join(relToRoot, postCSSName))
	}

	header := rewriteIndexLinks(tmpl.PostHeader, relToRoot)
	footer := rewriteIndexLinks(tmpl.PostFooter, relToRoot)

	page := fmt.Sprintf(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>%s</title>
  <link rel="stylesheet" href="%s">
</head>
<body>
%s
%s
%s
</body>
</html>
`, html.EscapeString(p.Title), cssHref, header, rendered.String(), footer)

	postPath := filepath.Join(postDir, indexFileName)
	if err := os.WriteFile(postPath, []byte(page), 0o644); err != nil {
		return fmt.Errorf("write post page %s: %w", postPath, err)
	}

	return nil
}

func rewriteIndexLinks(content string, relToRoot string) string {
	if relToRoot == "." {
		return content
	}
	target := filepath.ToSlash(filepath.Join(relToRoot, indexFileName))
	content = strings.ReplaceAll(content, `href="index.html"`, fmt.Sprintf(`href="%s"`, target))
	content = strings.ReplaceAll(content, `href='index.html'`, fmt.Sprintf(`href='%s'`, target))
	return content
}

func copyReferencedImages(markdown string, srcBase string, destBase string) error {
	matches := mdLinkPattern.FindAllStringSubmatch(markdown, -1)
	seen := map[string]struct{}{}

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		normalized := normalizeLinkTarget(match[1])
		if normalized == "" || !isLocalImagePath(normalized) {
			continue
		}
		if _, ok := seen[normalized]; ok {
			continue
		}
		seen[normalized] = struct{}{}

		sourcePath := filepath.Join(srcBase, filepath.FromSlash(normalized))
		if err := validateRelativePath(normalized); err != nil {
			continue
		}
		if _, err := os.Stat(sourcePath); err != nil {
			continue
		}

		destPath := filepath.Join(destBase, filepath.FromSlash(normalized))
		if err := copyFile(sourcePath, destPath); err != nil {
			return err
		}
	}

	return nil
}

func normalizeLinkTarget(raw string) string {
	target := strings.TrimSpace(raw)
	if target == "" {
		return ""
	}

	parts := strings.Fields(target)
	if len(parts) == 0 {
		return ""
	}
	target = parts[0]
	target = strings.TrimPrefix(target, "<")
	target = strings.TrimSuffix(target, ">")

	if idx := strings.IndexAny(target, "?#"); idx >= 0 {
		target = target[:idx]
	}

	lower := strings.ToLower(target)
	if strings.HasPrefix(lower, "http://") ||
		strings.HasPrefix(lower, "https://") ||
		strings.HasPrefix(lower, "data:") ||
		strings.HasPrefix(lower, "mailto:") ||
		strings.HasPrefix(lower, "#") ||
		strings.HasPrefix(lower, "/") {
		return ""
	}

	return target
}

func validateRelativePath(path string) error {
	clean := filepath.Clean(filepath.FromSlash(path))
	if clean == "." || clean == ".." {
		return errors.New("invalid image path")
	}
	if strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return errors.New("image path escapes source")
	}
	return nil
}

func isLocalImagePath(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".svg", ".webp", ".bmp", ".ico", ".avif":
		return true
	default:
		return false
	}
}

func sortPosts(posts []post) {
	sort.Slice(posts, func(i, j int) bool {
		if !posts[i].Date.Equal(posts[j].Date) {
			return posts[i].Date.After(posts[j].Date)
		}
		return posts[i].Title > posts[j].Title
	})
}

func renderIndex(posts []post, tmpl templates) error {
	var links strings.Builder
	links.WriteString("<ul>\n")
	for _, p := range posts {
		label := p.Title
		if p.IsDraft {
			label = "[DRAFT] " + label
		}
		href := filepath.ToSlash(filepath.Join(p.Slug, ""))
		links.WriteString(fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", href, html.EscapeString(label)))
	}
	links.WriteString("</ul>\n")

	page := fmt.Sprintf(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Blog</title>
  <link rel="stylesheet" href="%s">
</head>
<body>
%s
%s
%s
</body>
</html>
`, indexCSSName, tmpl.IndexHeader, links.String(), tmpl.IndexFooter)

	indexPath := filepath.Join(outputDir, indexFileName)
	if err := os.WriteFile(indexPath, []byte(page), 0o644); err != nil {
		return fmt.Errorf("write index page: %w", err)
	}
	return nil
}

func copyFile(src string, dst string) error {
	input, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open %s: %w", src, err)
	}
	defer input.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("create directory for %s: %w", dst, err)
	}

	output, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create %s: %w", dst, err)
	}
	defer output.Close()

	if _, err := io.Copy(output, input); err != nil {
		return fmt.Errorf("copy %s to %s: %w", src, dst, err)
	}

	if err := output.Sync(); err != nil {
		return fmt.Errorf("flush %s: %w", dst, err)
	}

	return nil
}
