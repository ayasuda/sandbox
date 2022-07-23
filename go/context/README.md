Go context samples
=====

# prog1

the http server sample it is configured timeout by context.

```
$ go run prog1.go
all: baz, 
all: baz, foo, 
all: baz, foo, qux, 
all: baz, foo, qux, bar, 
Response: baz, foo, qux, bar,  (719.454247ms)
```

# prog2

the sample code for context.WithCancel.

```
$ go run prog2.go
start infLoop
do cancel
  exit infLoop
```
