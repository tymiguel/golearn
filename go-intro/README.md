# Introduction to Go

Go (golang) is a programming language developed by Google in 2009. It's similar to C except that it has garbage collection, memory safety, structural typing and CSP-style concurreny. It was built because the developers had a shared dislike for C++.

Go is:
- statically typed
- compiled

Go has:
- runtime effciency like C
- readability and usability like Python of JavaScript
- high performacne networking and multiprocessing

Go may appear like an OOP langauge but it is not. It doesn't not have type heriarchy. Methods in Go are more general than C++ or Java and can be defined for any sort of data, not restricted to structs (aka classes).

## Installing GO

- First download go binary release at: https://golang.org/dl/
- Set up local environment: https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go

## Setting up directory

When go compiles and intalls tools, it puths them in the $GOPATH/bin directory. Having a structure like this allows the project to be able to work with the `go get` tool.

A `go` workspace is a directory with `/bin` and `/src` at it's root. Most developers will keep all their code and dependencies in a single workspace.

The `go` tool builds and installs binaries to the `bin` directory.

Structure:
```
# $HOME=/home/tmiguel
# export GOPATH=$HOME/go
# export PATH=$PATH:$GOPATH/bin   
$HOME/go                                <- go directory, root workspace
$HOME/go/bin                            <- go binaries, executables, tools
$HOME/go/src                            <- go source code files
$HOME/go/src/github.com                 <- canonical packages
$HOME/go/src/github.com/tmiguel/golearn <- git project directory
```

The `src` subdirectory will contain the version control repositories such as `github.com`. When you import libs, you will see this prefix which allows for canonical (reference fully qualified package) imports of code.

Canonical import paths make it resilient to repo changes.

### Other important directory notes

`go build`  <- compile code to executable that will be named after the project which allows to be run from command line
`go install` <- creates executable and puts it in the `bin` folder; if there are dependencies, will cache them and put them in `pkg` folder

## Running go

Using the command from the CLI, use:

```
go run hello_world.go
```
## Resources

- Wikipedia: https://en.wikipedia.org/wiki/Go_(programming_language)
- Golang.org: https://golang.org/
