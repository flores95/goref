# go reference

This repository contains a set of services and a command line app that demonstrate features of the go language. You'll find references below on where different features of go are being used. You'll also find descriptions of the design/architectural patterns that are in use. Please see the READMEs in the **apps**, **services**, and **frameworks** directories for descriptions of their functionality.

---

## __go features__

### structs
> Structs are everywhere in this reference repository. Just search for `struct` and you'll find tons of examples. Structs are go's _classes_ or the closest we'll get to a class. It's important to remember that the functions and attributes connected to a struct that start with a **lower-case** letter are private (not _exported_ in go terminology) and those that start with a **Upper-case** letter are exported (you can think of them as public :) ). This is really important when using the JSON libraries to Marshal/Unmarshal structs. Only exported attributes are converted. There are workarounds. Take a look at `apps/models/user.go` for an example.

### functions

### custom types
> Custom types can be super powerful in go. They allow you to build some business rules into types that are validated at compile time. That's way before a user sees your code and that's a good thing. The logging framework has the best example. Take a look at `fameworks/log/logger.go` and the `Level` type for an example. 

### interfaces

### go routines

### packages

### unit testing / mocking

---

## __design patterns__

### dependency injection

### factories

### inheritance (or as close as it gets in go)

