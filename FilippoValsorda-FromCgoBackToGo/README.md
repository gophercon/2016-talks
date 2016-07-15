## From cgo back to Go

Filippo Valsorda / [@FiloSottile](https://twitter.com/FiloSottile)

As soon as you `import "C"`, a lot of the things we love about Go are gone.

This talk helps you walk your way back up from the bottom of that pit.

Topics: 

* Using C memory to avoid the GC
* Callbacks with user data 
* Profiling
* Avoiding the cgo function call cost
* Using incremental builds to "make compilation fast again"
* Vendoring C libraries
* Building static binaries
* Getting `go get` to work, with sacrifices
* Even cross compiling, if you are brave

The slide on callback handles comes with a bugfix delivered at Q&A time by Ian Lance Taylor.

[View the slides online](https://speakerdeck.com/filosottile/from-cgo-back-to-go-gophercon-2016).
