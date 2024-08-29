# learn-golang

## Go Resources & Books

- [A Tour of Go][def8]
- [Go By Example][def9]
- [Learn Go with Tests][def10]
- [🌻 Effective Go][def11]

## Topics with Code Examples

### Data Types

- [Numbers, Strings, Boolean][def3]
- [Arrays, Slices][def16]
- [How Capacity increases on slice.append()][def16]
- [Slices Deep vs Shallow Copy][def17]

### Maps & Structs

- [Maps & Structs][def19]
- [Maps are not Reference Variables][def18]

### Functions, Closures, Flow Control(loops, if-else)

- [Functions, Anonymous Fns, Recursive Fns][def2]
- [Variadic Functions - Ever wondered how `fmt.Println` accepts any number of arguments without passing them as a list?][def22]
- [Closures in Go][def]
- [Classic for loop, for-in, for loop as while loop, Range over integers][def15]

### Pointers

- [Pointers Introduction][def12]
- [🌻 new() vs make() & Zero-Values of all types][def20]
- [🌻 Does Go supports Pass By Ref ?][def14]
- [Using Pointers with Structs][def13]

### Composition & Embedding

- [Go doesn't support Inheritance. How can we achieve the similar functionality using the concepts of Composition & Embedding][def21]

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
[def12]: basics/08Pointers/pointers.go
[def13]: basics/08Pointers/pointers_with_struct.go
[def14]: basics/08Pointers/pass_by_val_vs_ref.go
[def15]: basics/01FunctionsAndFlowControl/flow_control.go
[def16]: basics/04ArraysAndSlices/arrays_slices.go
[def17]: basics/04ArraysAndSlices/slices_deep_copy_shallow_copy.go
[def18]: basics/05StructsAndMaps/maps_are_not_ref_variables.go
[def19]: basics/05StructsAndMaps/structs_maps.go
[def20]: basics/08Pointers/new_vs_make.go
[def21]: basics/09Composition&Embedding/README.md
[def22]: basics/01FunctionsAndFlowControl/variadic_functions.go
