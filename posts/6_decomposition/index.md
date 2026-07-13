title: decomposition and long test times
date: 2026-06

## Introduction

I have been following a pattern for some time, riffing on the old `qtest` approach of client/server

- single binary
- client and server
- client is CLI and TUI
- server is headless openapi AND embedded site
- server batteries included (registration, email, site, api, metrics, database)
- ghcr.io image
- brew installable

Where the point is convenience.  It comes at a complexity cost on the side of the developer but that's
sort of my default place - engineer takes the pain so the user doesn't have to.  It "just works".

OK so now you roll forward agentically and the idea is all too real in most all my apps.  And now I get
to the point where I want red/green testing and full test passes across the whole estate, combined with
"ease".  This is where I find conflict in *Time*.

My tests are slow again.  Unit is nippy, integration - especially integration around anything UX-ey, is s-l-o-w. 

So now I want to decompose, but rethink what that is.  Different components that contribute to an overall binary - therefore I think I want *hard contracts*

The agent needs to then
- live in a monorepo
- touch only a subset
- and udnerstand the api boundaries that indicate the scope of testing required

OR

I have many small tools which counters the "monorepo single binary".  What to do.




