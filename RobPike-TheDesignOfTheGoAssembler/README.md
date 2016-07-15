## The Design of the Go Assembler

Rob Pike, July 2016

[@robpike](https://twitter.com/rob_pike)

The Go assembler was rewritten for the 1.5 release, replacing a suite
of C programs, one per architecture, with a single binary, written in
Go, that supports all architectures. The usual variables, GOOS and
GOARCH, are sufficient to configure it for any environment.

This talk will explain how this extreme portability is achieved for
such a non-portable problem. The answer lies in the structure and
origin of the Go compilation tool chain, a mostly machine-independent
input language, and a lot of automation. Even for non-assembling
Gophers, there are lessons in the design.

[View the slides on line](https://talks.golang.org/2016/asm.slide)

