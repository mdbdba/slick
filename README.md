# SLIck (SLI check)
In thinking about how to make it easier for developers to define SLIs that reflect a user's or consumer's experience, there should be a tool that can generate metrics for a service that the service itself cannot provide.  A point of view outside of the service that can be evaluated to determine if it is doing what it was designed to do.

From an Ops perspective, this would also be helpful because we could set up simple sanity tests for *aaS providers (put/get records into Kafka, ES, DBs, Redis)

The thought behind Slick is to provide a way to interrogate services with definitions of checks (simple url args or JSON objects) and generate a set of metrics.   Idea being to provide a simple way to get some reliability feel for services that devs can use for SLO definition

Use cases:
Test service reliability with a check record.  e.g.  roller/?roll=5d1 eq 5
ES and Kafka add record / simple read record test

