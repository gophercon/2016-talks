# Go Without the Operating System

Bernerd Schaefer 12 July 2016

Go lets you to compile your program into stand-alone binaries which can be
safely shared across systems without worrying about a missing dependency. But
there’s one thing even a stand-alone executable depends on: the operating
system. What if you could remove that dependency, too?

[AtmanOS] is a project built to explore that question.

It allows you to compile ordinary Go programs and run them on cloud providers
like Amazon’s EC2, without a traditional operating system.

[View the slides online][slides]

  [AtmanOS]: http://atmanos.org/
  [slides]: http://atmanos.org/talks/go-without-the-os-gophercon-2016/
