## Structured Logging and You

Antoine Grondin / [@AntoineGrondin](https://twitter.com/AntoineGrondin)

You’re about to log, you have lots of context in your program:

- variable names
- values

It’s all well organized into your process’ memory.  But then you take all that well organized stuff and mix it up in some printf’d together string and you lose all the nice order you had. Then you collect that mumbo jumbo for hours, from bunch of processes and machines, and eventually try to make sense of it with code, always trying to recover the nice “name -> value” you had in your program.

Maybe you should not have thrown away that key-value mapping?

[View the slides online](https://docs.google.com/presentation/d/1632rG5GJD-bGqzaV4ILBnHEJ2mnw8rVbulp-LyKTtuY).
