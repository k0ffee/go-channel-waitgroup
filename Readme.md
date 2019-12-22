Small example for using channels and waitgroups in Go.

For every incoming request, it fetches data from external
services concurrently.

Any responses are put into a channel and finally consumed from
there for presenting the result.
