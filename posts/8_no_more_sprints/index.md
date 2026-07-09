title: no more sprints
date: 2026-06-21

I want to test removing sprints altogether and think more about releases, where a release is a set of features, where a feature is a set of epics, where an epic is a set of stories.

this means the arrangement of stories woudl be

1. a backlog containing epics and stories
2. a story links to an epic
3. an epic links to a requirement

The reverse is therefore implied:

1. you start wtih a requirement
2. it is refined until it is consistent
3. it is borken down to epics and stories
4. they are verified as consistent between themselves and in pursuit of the requirement

This means a person -= the Product Owner - than then say: I want feature 1, 2 and 3 in the next release.

So you make a release candidate of features or requirements
    which may then contain dozens, hundreds of stories

The engineering agents then

1. picks up the stories in any order based on the prioritisation and dependency graph
2. performs the enginering necessary
3. marks them as complete (goes through a workflow cycle of all roles on each story during each phase)
4. the human then approves the feature as compliant: OR provides feedback on why they do not comply, which then means the requirement is re-done.

Note: there isn't really an engineering agent or a refinement agent etc. Really it's just an agent - which is provided motivation in the form of the prompt containing all the information necessary, including references to external tools, skills, e.g. mcps etc. depending on the context of the work being required to do.

Ok bak to removing sprints.

So the new "container entities" are

Release
    has a title and purpose
    contains features
    has a title and an aspirational date for delivery
    has a status of either in design, in progress or complete
    tracks dates of change of status

Feature
    This is a ticket of type feature which is the "grand plan"of what it is that is necesary.  This will
    be refined with a human and agent using teh feature ticket as well as the codebase and all project documentation.

These entities apply to a project.

The backlog entity then is the container of
    Feature
        Epic
            Story/Bug


At any point a user can decide to put the feature into a release, as long as the release is not yet active.

A user can "clone" a feature if they want to exten the functionality of it.

---

the above I want you to work on

UX
CLI
Data Model
Documentation

in a new branch, feature/kanban

I want you to get it all working and update the tk demo to populate using that new entity model.



(IN THE CORPORATE WORLD)

Sprints, Scrum, Agile, aka "That's not Agile"

(Please leave me alone if. you are going to provide a clarification on what velocity is, or say fibonacci).

We used sprints to know 
    - where we "are" - e.g. sprint 16 
    - estimate how long something might take (that'll take 3 sprints)
    - organise a motley crew of technology people with a blunt instrument and standard process

We measure velocity - how much work gets done in a period of time with a fixed capacity - to work out if we are going faster or going slower - meaning more work getting done relative to the last sprint.

Work was measured in an abstraction unit called a story point.   We set those points using a game called
planning poker where we estimated the complexity of the feature relative to other things.  For example, this thing is small, 1.  This thing is twice as large as that small thing, 2, this thing is mega, 10 times larger than the 2 thing, 20.   Then we would look back on our previous sprint and say: we tend to average out 6 of those abstraction points per person per sprint.  Ok what's our capacctiy.   We tend to reckon we can do like 1000 things but actually achieve 12.  Ok maybe lets set a sprint to contain 12 things.

And then we ask questions in ceremonies called retrospectives: what went well, what can we improve, who wants a donut.   And then we repeat.   Welcome to my life for the last 19 years.  

I quit that life!


------

No more sprints.  just releases.   We agree all the things that need to be done, then we wrap them into a bag of featues, then we tell the agents to implment them all.   We guide, shepherd and review, but we dispose of the sprint.   It has become overhead that is no longer necessary.

What is necessary is

    - product owner involvement in prioritisation
    - rigor in the requirment definition



Assumptions

- CI/CD implying repeatable, consistent, predictable and reliable deliveries
- no sacrifice in quality owing to a robust test framework and culture
 
 - At any moment we should be able to create a release of features
- We then elect those moments with rules (never on a friday, 2 weeks cooling down)
- We then optimise those rules to go faster


