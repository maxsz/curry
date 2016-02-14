> I like my chicken like my functions of multiple arguments: curried. - Michael Snoyman

curry
=======

Function curring for [Go](https://golang.org). 

Usage
-----

1. Install `go get github.com/maxsz/curry`
2. Add `//go:generate curry` as a first line of any go file within a project.
3. Run `go generate` before `go build`
4. Use generated curried versions of any functions in your code. Example:

```go
// Example function definition
func example(first, second string) { return first + " " + second }

// Use generated, curried version of the function
hello := exampleC("hello")
hello("world")
```

Motivation
----------

Although I really like go, I think it's biggest draw-back is the missing
generics support. I wanted to explore and learn about the go type system and
code generation tools and how they could help mitigate the missing generics
support. This is what came out of it.

References
----------
[The Go Blog - Generating code](https://blog.golang.org/generate)
