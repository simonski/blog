TODO

Upgrade this to a go binary installable "blog" tool

# setup a. git repo
blog init

# list all posts
blog ls

# publish or unpublish
blog publish / unpublish

# metrics
blog metrics -id X

# new entry
blog new "foo"
makes a folder, marks the state machine that foo is new and in progress

# static site server
blog server/serve

# build static site under dist/static
blog build

---

design

a small database that holds metadata (sqlite)
plaintext, markdown, html, you-name-it under the post itself
an index that is maintained and rebuilt

