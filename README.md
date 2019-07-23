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
>Routines in go provide the ability to execute in parallel or on multiple threads. The only rudimentary example of this in this repository is in `apps/cli/app.go`. The **RunAsync** function will run each **Processor** in it's own thread asyncrhonously. 

### packages

### unit testing / mocking

---

## __design patterns__

### dependency injection

### factories
> Since go doesn't have the concept of a constructor for structs, you'll find the use of struct/class factories all over go code. This usually takes the form of a function called  _NewStructName()_. You can find many examples of this patter in this repository. 

> More elobarate examples of factories that generate interfaces based on configuration data can be found in `frameworks/log/logger.go` and `frameworks/storage/store.go`. Each of these **new** functions will generate an implementation of the interface based on configurator settings. 

### inheritance (or as close as it gets in go)
> Inheritance is not a built in go construct. The concept of inheritance (inheriting behavior of a parent class and chosing to use it or extend it) is easily achievable. If two structs implement the same interface, one can become a base/parent/super class of the other by simply adding it to the child struct. You can find an examples of this patern in this repository. 

>The **Configurator** interface. `frameworks/config/dotenv.go` implements the **Configurator** interface and uses `frameworks/config/base.go` as it's "parent" or super class. The reference in **dotenv** to **base** is called super but that's just to make old school OO developers happy. There's nothing special about the attribute name "super".
