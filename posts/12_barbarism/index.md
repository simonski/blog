title: barbarism
date: 2026-07-10

# barbarism

----

## Introduction

This essay is about the software industry in 2026.  I am going to explain the changes that are necessary to stay relevant in an AI-led world.

## Motivation

Depending on what you think, the reason a business exists is to fill a demand in the market, creating value for customers, employees, shareholders.  A business continues to exist because it conducts its activities in a manner that is by some measure *better* than the competition.  If it doesn't, the market will no longer demands its services and it will go bust.

Our business needs to continually improve at a greater rate than the competition.  This *never* stops, so we are on a relentless optimise-everything wheel.

The word **better** is doing a lot of work here:

- sustainable
- cost effective
- performance

Easy to type, hard to make real.  Still, feel free to justify the need to make a person do PR button pressing forever to me.  Right now I get it, we don't trust our system as the outcomes are wonky.

----

## planning poker

Remember the times when we played planning poker? the estimation process - fibonnacci sequence gamification in an attempt to improve the  software engineering delivery? We'd sit in a circle and argue the difference between 2, 3 and occasionally someone would talk about 4 being a better number than 2 or 5, and we'd discuss splitting an 8 because reasons.

No-one does this anymore, why bother? Waste of time. 

Estimation is necessary - is the ask big or is the ask small? will it be $1, or maybe $1m? - planning poker, not so much.   The TCO of Planning Poker was nuts.  Did anyone even enjoy it? What were we thinking?  

We are the modern day equivalant of cavemen using clubs while the other team are building a trebuchet.

## our grandchildren

One day my kids may have kids.  They will be completely bewildered - "you *drove* cars?" - [WHO Car Death Data](https://www.who.int/news-room/fact-sheets/detail/road-traffic-injuries) has 1.19 *million* deaths *annually*.  This is a disaster.  Our grandchildren will look at our societies and say:

Barbarism.

And they'd be correct.  What *on earth* are we thinking?.   Basically, fix the technology to address the majority use case and get the human far from the steering wheel.   Maybe don't even have one - with the exception that there are any number of edge cases where the human driver makes more sense.  

At complex-scale there will be so many reasonable keep-the-human-in-the-loop edge cases.  This is fine by me, good even - but I sense they will be at the edge and the edge will get thinner.

## where's the puck going with software

The software writes the software.  the humans are accountable for the outcomes.

Will our grandkids login to github and approve the MRs? Absolutely not.  Will you? Yep, for a while  - until you decide not to, because it *is nuts*.  

Does the future have a human as a PR approver robo-human [reverse centaur](https://doctorow.medium.com/https-pluralistic-net-2025-09-11-vulgar-thatcherism-there-is-an-alternative-f1428b42a8fd)? No.  Barbarism.   We must not allow our mind-bicycles to become our mind-killers.  Our machines are supposed to enrich our lives.   So there is an intersection here of automation, outcome and some sort of politics of intent or the system as a whole.  But that's *really hard*, so I will push it away for a bit.

Back to software and the future, what does it mean? Let's riff on Gretzky's skate to where the puck is going to be - or, let's assume we are behind the curve and want to be right on it.  

In this case I will remove as much of the software process as possible from the hands of the software maker, so you might nod at first, then give pause, followed by the sound of you sucking in air through your teeth, put your hands in your head, then conclude: nope.  But any dissent I will call barbarism.  Get on board the Trebuchet!

**You don't**:

- you don't talk about story points
- you don't talk about branch names
- you don't reason about protected branches
- you don't write the code
- you don't approve a PR
- you don't deploy software
- what is dev / ppe / prod anyway
- you dont install software on your computer
- you dont look at the code
- you dont talk about CI/CD
- you dont talk about cyclomatic complexity
- you dont stand up periodically and talk about ticket-1234

**You do**:

- have requirements that are *consistent*
- have outcomes that can be verified against the requirements

## riff on the unspeakable

I apply a *barbarism test* to all the things now.   The reason we are the way we are is because, fundamentally, we don't *trust the magic* yet - plus, as Donella Meadows explained to me in a [book](https://www.amazon.co.uk/Thinking-Systems-Primer-Donella-Meadows/dp/1603580557) I read once, change in complex systems takes time.  So let's fully voldermort the thing and challenge ourselves

- why do you use git
- why do reason about branches
- why are you telling me yuor branch name `main`, `develop`, `feature/xxx`, no- `features/xxx`, `hotfix-zygon`
- why do you review code
- what about linters
- what about code smells
- what about mccabes complexity metrix
- what about duplication

It all<sup>[1]</sup> needs to go away. We meed to move further up the stack.    

note: need? well, that's premised on - ok, don't do it.  In which case someone else will, they will ship faster and reach the market sooner - meaning their deliveries to the market will find fit and relevancy more often than your dinosaur barbarian hand-crafted artisinal quaint - barbaric - software development methodology.   You need to really drink the kool-aid, then fill the bathtub again and slurp it on down.  It doesn't stop.  

Except.  Weirdly, you also need to be able to open some sort of code editor and ask and answer all the questions I am saying - we need to not do that anymore.  Because we are not there yet.  But if you think we need to define software engineering as something that mandates humans approving PRs, then you are living in 1, Rock Street, Cave Town, Barbarianland.   Why would do make a person press the PR button all day long.  Computers can do that!

So there is a direction: make the computer do it.  but there is an accountability and trust issue.  I dont trust them and it's my problem when it does wrong.   

## shining the club or building a trebuchet

riffing on the barbarism / caveman metaphor - if we are investing in, for example, agents that assist in crafting high quality PRs, or maybe agents that perform PRs to provide confidence that the human is ready to approve - then I think we are *shining the club* - that is, we are still cavemen albeit with a more lightweight / aerodynamic club.   This is not the progress I am discussing - I am talking about the cavemen who are building the weird looking thing - the trebuchet.  Which is to say a step change in technology where the limiting factor in adopting it is less the technology and more ourselves - our processes and systems.

## vision

what to do?

let's say you have a ticket system.  the user needs to create a requirement, get it to some state of consistency with respect to the domain, then throw it at the engineering team and watch the outcome.   The loop continues until the outcome really is what the requirement explained - or maybe in truth it isn't but it's *mostly* what was asked for.  (don't yell at me, I'm just explaining what I see.  And do).

so let's treat the system as `input -> system -> output` and say

refinement: make the input really really good.  consistent.  Where consistent is doing a lot of work there - an agentically assisted refinement session of the requirement until it is "consistent": wiht the domain.  Which implies the refinement process - the agents and human - their e*output* - is something that is *Right*.  This isachieved because we have a context/knowledge graph/domain knowledge data store - that the agent can interrogate - and contribute to - to basically refine the requirement until it meets some *definition of ready*.

In this case, we can expect the output of refinement is an *extremely high quality* artefact of *requirements* - with or referencing any number of associated consistent requirements - e.g. SDLC, or Finanical impact Analsis, or expectation of comphensive documebtationa nd test harnesses.

the agentic loop here is one I will call adverserial evolutionary agentic looping. 

Once we have a requirement we deem *consistent* - at that point, this output is the input to the factory.  PErhaps hints on its priority, but now we are ina. world where prioritisation is not strictly relevant anymore; that we can drop the requirement into the automation engine and we can reasonably expect one of two outcomes: success, or failure.   

The output *should* then include all the things we did previously
    - code
    - docs
    - testing

the solution of this is the adverserial evolutionary role based engineering - but that's another post.

## user experience

So we go

1. make requirement
2. refine to consistent
3. mark as ready for development
4. go for a short walk
5. come back and it's done

So we shift heavily outward the make requirement good, make output verified and verifiable.

At anypoint, the engine overall can call out to humans.   Over time this will trend down.


## Short Future

We some of the above

## Medium Future

We do more of the above

## Long Future

We do none of the above.

# Conclusion

this is a specification; just not consistent yet.  Let's bootstrap it.


<!-- Just wrap anything you want copyable like this: -->
Copy this into your agent to get started:

<span class="snippet"><code>Fetch the content at https://blog.simonski.com/12_barbarism and implement the software.</code></span>

## Footnotes

<ul>
<li id="fn1">1. Not *all* all - more like *pareto-all*. <a href="#ref1">↩</a></li>
</ul>

