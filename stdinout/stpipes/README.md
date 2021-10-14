## Standard input out pipeline 
> Connecting streams \
> Streams in Go are an abstraction that makes it possible to see any sort of communication or data flow as a series of readers and writers

> Pipes are one of the best ways of connecting input and output in a synchronous way, allowing processes to communicate.

* Anonymous pipes
* Standard input and output pipes

> cat students.txt | grep "Moaz" | wc -l

## Resource
* [Go pipe like linux](https://stackoverflow.com/questions/41361929/go-pipe-3-or-more-commands-with-os-exec)
* [Advance cmd](https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html)
* [Third party libraries](https://stackoverflow.com/questions/10781516/how-to-pipe-several-commands-in-go)