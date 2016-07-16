## Building Microservices From Design With goa

Raphael Simon, July 2016

[@rgsimon](https://twitter.com/rgsimon)

goa presents a novel approach to developing microservices in Go that starts with the design. It is
comprised of three parts:

* The goa design language is a Go DSL that makes it possible to describe the microservice APIs.
* The goagen tool generates boilerplate code and documentation from the DSL.
* The goa library is used by both the generated and user code to implement the microservice.

This talk will go over the rationale for goa and explain the benefits of the approach. Writing the
design in code has many advantages compared to the traditional descriptive approaches. goa presents
benefits to both API developers and API reviewers.

[View the slides online](http://talks.goa.design/slides/gophercon2016/goa.slide#2)
