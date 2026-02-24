PROMPTS

create a makefile with targets
    build:
        runs a go script to create the site
    run:
        serves the output directory on localhost:8000
    deploy:
        deploys the content using rsync to the root of blog.simonski.com

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

