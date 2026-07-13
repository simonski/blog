#!/bin/sh
set -e
rsync -avz --delete output/ blog.simonski.com:blog/
