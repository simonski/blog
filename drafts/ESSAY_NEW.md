# The Software Factory

*2026-02-16*

Can a system of agents replace the traditional software development lifecycle? Should it?

The Software Factory (TSF) is an approach to answer this question. The goal is to automate and augment the current SDLC process where we expect to improve productivity. The desired outcomes are:

- Programmers operate at a higher level of abstraction
- The number of programmers in the loop can be reduced for the same or greater amounts of work
- The number of agents in the loop is increased
- The quality of work is maintained or improves
- Further understanding of agentic engineering practices
- Data to understand and represent cost-value in software engineering

The approach is to model the roles and processes we have now and automate them. Applying this across the software engineering stack is the experiment.

Three key challenges frame the work:

- **Trust** - How do we establish trust if the human is taken out of the loop? How do we trust a computer will both create the code and tell us that the code is correct?
- **Repeatability** - How can we demonstrate a given requirement will result in an expected outcome, repeatably?
- **Provability** - How can we "show our workings" to demonstrate WHY a given outcome occurred?

What follows is a tour of the factory - each component introduced one at a time, building up to the full system.

---

## 1. The Linear Mental Model

I *thought* that software went from an IDEA through some PROCESS to REALITY. Zooming in one level of abstraction, we apply different roles, lifecycles and methods. For example, with Roles:

IDEA -> Product Manager -> Product Owner -> Business Analyst -> Architect -> Tech Lead -> Engineer -> QA -> Release Manager -> REALITY

![The Linear Pipeline](diagrams/01-linear-pipeline.svg)

The job is like a baton, handed on to the next in line.

Obviously this is **not true** - jobs were handed back and forth - however this was my original mental model.

---

## 2. The Nonlinear Reality

In fact what happens is everyone talks to everyone - sometimes. The Engineer clarifies a point with the QA who talks to the BA and the Release Manager coordinates with the Engineering Director. And so on.

![The Nonlinear Reality](diagrams/02-nonlinear-reality.svg)

Software development is nonlinear. This is the first insight that shapes the factory.

---

## 3. Role

All work is carried out by a discretely named ROLE - for example a Business Analyst.

But there is a conundrum. In a human team, the human IS the role - meaning they are not differentiated effectively. "Tribal knowledge" remains in the head of a specific human. There is no such thing as a "QA Engineer" - rather there is a "QA Engineer called Sally".

Sally has years of experience and value. Sally is the memory. Sally "knows" the different roles to assume or to go find in the company.

In an automated system, there can be no place for tribal knowledge. This "institutional" memory needs to be modelled. A ROLE contains as much of the institutional memory as possible - it describes the job function, motivation, objectives, boundaries and relationships between other roles.

![Introducing: The Role](diagrams/03-the-role.svg)

The outcome of work by a Role is then received by other Roles with different motivations - those roles will implicitly or explicitly appraise the quality of work, using or rejecting it.

---

## 4. Story

A STORY is the objective itself. It contains the objective, its history, which roles did what to it. The data model is significant as it will contain sections that are effectively a living document of change.

A Role works on a Story. The Story is the unit of work flowing through the system.

![Introducing: The Story](diagrams/04-the-story.svg)

For any piece of work, it will have lineage - the effort taken, by who and when, the outcome. This state represents the "history of the story" and can be useful informing future decisions.

---

## 5. Skills

SKILLs are reusable technology capabilities that a ROLE calls upon to achieve a given objective. A Business Analyst might use a "Requirements Elicitation" skill. An Engineer might use "Code Generation" or "Test Writing" skills.

Skills are separate from Roles because the same skill can be shared across roles, and roles can be composed with different skill sets.

![Introducing: Skills](diagrams/05-skills.svg)

The theory is the state - the memory - is the aggregation of ROLES, SKILLS and a STORY.

---

## 6. Council

As work is completed by a Role, the outcome will be assessed by a COUNCIL. This is effectively another role - a role "of" roles - which acts as the motivation to decide what is the appropriate next step.

It is NOT that the work is DONE. It is more that the work of the Role is finished *for now* and the decision to progress the work in any direction is the result of the council decision.

![Introducing: The Council](diagrams/06-the-council.svg)

So work looks like this:

```
ROLE1 + WORK -> COMPLETED -> COUNCIL -> OK    -> ROLE2 + WORK
ROLE1 + WORK -> COMPLETED -> COUNCIL -> NOT OK -> ROLE3 + WORK
```

Until the COUNCIL decides the work is complete. The council is an appropriate set of Roles in any given moment - it is not a fixed decision making process. Different roles and motivations will apply.

An initial WORKFLOW structure will be provided by the OPERATOR which will act as advice on the available roles and various acceptance criteria that must be met.

---

## 7. The Living Document

Every pass through a Role and Council adds to the Story. The Story is a living document - it accumulates state from every interaction.

![The Living Document](diagrams/07-living-document.svg)

Caution will be needed as there is no upper bound to the history of a given problem - so using contemporary LLMs we will need to address this with regards to relevant context sizes. A story that presents its own challenge may spawn another task to "prepare this story for the role" - a summarisation agent.

---

## 8. Consensus

How does software get built to any reasonable quality? Some roles get together in a little council and decide to agree or disagree - the definition of the job, the outcome of work being successful or not.

This happens in micro-meetings with perhaps some established processes or ceremonies - but it also happens adhoc in methods we have not properly captured.

The progression of any work is as a result of CONSENSUS. Where work proceeds with consensus of 1 (the engineer says: it works, trust me), is where there is high risk.

![Consensus](diagrams/08-consensus.svg)

Where there is NO consensus, escalation occurs. It is possible we will learn that some outcomes cannot be decided autonomously - the system could escalate to the HUMAN in the loop. This is not a failure - rather an outcome we will learn from. This escalation may inform higher quality role descriptions, yielding fewer escalations.

---

## 9. The Factory

Putting it all together. An Operator bootstraps the system with Role definitions, Skill descriptions, and Workflow structures. Stories flow through Roles, gated by Councils, accumulating history, until consensus determines them complete.

![The Full Factory](diagrams/09-full-factory.svg)

---

## Fallibility

Teams are fallible and inconsistent. A Product Owner contradicts a prior feature. An engineer has a bad day. Four team members meet adhoc, decide on an outcome, and don't rewrite the specification.

An automated system *should not* be inconsistent in this manner. When the modelling is insufficient, the system should call this out and stop, or do something useful about it.

LLMs are also fallible. An LLM will hallucinate and output poor quality. We are not attempting to change a foundational model - hallucinations are baked into the system of LLMs. We accept this as a risk and mitigate it via decomposition, council and consensus.

The purpose of the Council is precisely this: task outcomes that result in poor quality are both accepted as a risk and mitigated through the council themselves.

## LLMs are the Critical Dependency

Critical to the whole endeavour is the use of large language models. The decomposition of work into Roles, Skills, Stories and the nonlinearity of the work means we expect:

- High token usage (back-and-forth)
- A ratio of consensus:code to be high:low (delivers trusted outcomes)
- Many roles and role compositions

The use of LLMs ARE the whole decision making process. This endeavour is about strictly controlling the entire context for any given prompt by modelling roles, skills, and stories.

## Cost

Removing linearity creates the risk of cost increase. To mitigate this we will introduce a circuit breaker and cost agent who will be part of the council.

## What is Success?

The outcomes will be directional:

- Can a requirement go from input to working code without human intervention mid-process?
- Does the council mechanism catch defects that a single-agent system would miss?
- What is the token cost per story point compared to a human team's cost?
- Does the living document provide sufficient auditability that a human reviewer can understand why decisions were made?

## Evolution

An SDLC process is not cast in stone - it changes. We will model this adaptive system with:

- A ticket tracking system
- The history of the work
- Separate roles with their objectives

In V1, we won't address evolution but prepare for it. Once we have built up a corpus of tasks and their outcomes, we will use this as learning material to see if we can create new Role descriptions or Skills that yield better outcomes.

## Conclusion

This document proposed an approach to automate the SDLC. The challenges of Trust, Repeatability and Provability are addressed through the decomposition of Roles, Stories and Councils.

The first version, tsf-1, will attempt to implement a reference factory and stop as it arrives at evolutionary roles. This means tsf-1 will be:

- A ticket engine
- Tickets are living documents
- A CLI to interact
- A website to interface and view work
- A fleet of workers

---

## Notes on Bootstrapping

In the first iterations of TSF, a HUMAN operator will need to describe the Roles, Skills, conditions under which a Council operates, so as to bootstrap the system - as well as provide requirements.

This means the first version of the system will need simple tools (CLI, Website) to allow an operator to define, refine and experiment.

Consensus loosely can be described as "all Roles in the current state of the work accept the outcome meets their objectives". This assumes each Role is of a high enough quality that we are not collapsing in a house of cards of false positives. Hence iteration in bootstrapping.
