PROMPTS

update the go and makefile to create a binary blogtool 'blog' with the following commands

    blog
        shows usage in multiline format

    blog init
        creates a new "blog" folder containing files, fails if folder already exists

    blog server
        serves static content from current directory

    blog post
        create new blog post

    blog idea
        create new idea post

    blog ls
        list posts and ideas sorted by most recent
            id, type, updated, title

    blog edit N
        opens vscode on the source content

    blog post "the title"
        create new post "the title"; fail if already exists, case-insensitive
        this will translate to
            posts/{id}_{compressed_title}/<template-files-token-switched>

    blog idea "the title"
        create new idea "the title"; fail if already exists, case-insensitive
        this will translate to
            ideas/{id}_{compressed_title}/<template-files-token-switched>

ID is an incrementing integer

split the blog into two areas - blogposts and ideas
    
    blog/posts
    blog/ideas

templates
    update to contain templates by type including site templates
        templates/site
        templates/posts
        templates/ideas

create a makefile with targets
    build:
        runs a go script to create the site
    run:
        serves the output directory on localhost:8000
    deploy:
        deploys the content using rsync to the root of blog.simonski.com


create an atom/rss feed for all content, update the makefiç

make <default> is "build run"

go: create a go script "blog.go" which
    - lives in the root
    - crawls the posts dir
    - generates html pages for all posts
    - updates an index.html page to link to all posts
    - writes all output to "output"
    - if the markdown contains links to images, copies the images to the correct output

crawling the posts
    for each directory under posts, there will be one markdown file

the file is a post in the format:
title: free form short title
date: yyyy-mm-dd
-------------------------------------------------------------------------------
BODY


This will be converted to html, retaining only the BODY, prefixing with the prefix template and suffixing with the suffix template.

The posts are ordered using the date descending then the title descenring and put into a list in the index.html.  The link text is the title: key value and the link URL should be the post path name.

index.html will refer to index.css
all posts will refer to post.css

index will pre fix with the index_header, index will suffix with the index_footer.

If --draft is passed in the call, then also include the drafts folder for posts and prefix the link with [DRAFT]

