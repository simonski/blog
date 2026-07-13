title: AI Providers
date: 2026-06-21

The AI Providers exist to represent "usage of a given model in a given provider". this means they
need to be created and store e.g. endpoints, identities, API keys.

Once the AI Provider is created and validated - that is, it should be possible to say "hello" and get a response back, it can then be makred as "ready" and then it can be assigned to a given project.

An AI Provider can be copied - so that the configuration is re-used but a different API KEY is attached and therefore the billing is directed to different areas.

A project or programme can then have multiple AI Providers "Attached" to them.

ALL usage of an AI Provider -that is, invoking an agent, is stored so that effective telemetry and usage and time and tokens can be tracked.

The UX should contain enough to configure and create these AI Providers.   

A user should be able to provide their own API KEY to a specific provider - in which case the AI provider values are "overridden"by the uer API KEY.  in this case the telemetry will then be visible to the user - as they are providing the KEY.  

This is also true of a team - a user can assign an API key to a team they are in or a project they are in.



