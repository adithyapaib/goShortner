# Go example projects

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

## Clone the project

```
$ git clone https://github.com/adithyapaib/goShortner
$ cd goShortner
```
https://go.googlesource.com/example is the canonical Git repository.
It is mirrored at https://github.com/golang/example.
## [hello](hello/) and [stringutil](stringutil/)

```
$ cd hello
$ go build
```
A trivial "Hello, world" program that uses a stringutil package.

Command [hello](hello/) covers:

* The basic form of an executable command
* Importing packages (from the standard library and the local repository)
* Printing strings ([fmt](//golang.org/pkg/fmt/))

Library [stringutil](stringutil/) covers:

* The basic form of a library
* Conversion between string and []rune
* Table-driven unit tests ([testing](//golang.org/pkg/testing/))

## [outyet](outyet/)

```
$ cd outyet
$ go build
```
