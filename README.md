# learn-golang

## Go Resources & Books

- [A Tour of GO][def8]
- [Go By Example][def9]
- [Learn Go with Tests][def10]
- [Effective GO][def11]

## Topics with Code Examples

### Data Types

- [Numbers, Strings, Boolean][def3]
- [Arrays, Slices][def16]
- [How Capacity increases on slice.append()][def16]
- [Slices Deep vs Shallow Copy][def17]

### Functions, Closures, Flow Control(loops, if-else)

- [Functions, Anonymous Fns, Recursive Fns, Variadic Fns][def2]
- [Closures in Go][def]
- [Classic for loop, for-in, for loop as while loop, Range over integers][def15]

### Pointers

- [Pointers, new() vs make()][def12]
- [Does Go supports Pass By Ref ?][def14]
- [Using Pointers with Structs][def13]

### Concurrency

- [Sequential vs Concurrency Execution - Fibonacci][def4]
-

### Dealing with JSON

- [Encoding & Decoding JSON][def5]

### Web Server in Go

- [Creating a simple Web Server using Gorilla Mux][def6]
- [Creating GET, POST, PUT, DELETE HTTP Methods - CRUD][def7]

### Context in Go

- [Context Package](basics/12Context/)

### SOLID Principles

- [**S**ingle Responsibility Principle](SOLID-Principles)
- [**O**pen Closed Principle](SOLID-Principles)
- [**L**iskov Substitution Principle](SOLID-Principles)
- [**I**nterface Seggregation Principle](SOLID-Principles)
- [**D**ependency Inversion Principle](SOLID-Principles)

[def]: basics/03Closures
[def2]: basics/01FunctionsAndFlowControl/functions.go
[def3]: basics/00DataTypes/datatypes.go
[def4]: basics/10Concurrency/04ConcurrencyParallelism/con_vs_parallel.go
[def5]: advanced/00-json/json.go
[def6]: advanced/01-simple-web-server/web_server.go
[def7]: advanced/02-go-mux-pro
[def8]: https://go.dev/tour/welcome/1
[def9]: https://gobyexample.com
[def10]: https://quii.gitbook.io/learn-go-with-tests
[def11]: https://go.dev/doc/effective_go
[def12]: basics/08Pointers/pointers_with_struct.go
[def13]: basics/08Pointers/pointers_with_struct.go
[def14]: basics/08Pointers/pass_by_val_vs_ref.go
[def15]: basics/01FunctionsAndFlowControl/flow_control.go
[def16]: basics/04ArraysAndSlices/arrays_slices.go
[def17]: basics/04ArraysAndSlices/slices_deep_copy_shallow_copy.go
