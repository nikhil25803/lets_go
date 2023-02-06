### Go lang

*Go lang* is a compiled language, but it can run file directly, without VM. Similarity with lots of languages like C, Java and Pascal.

It is not an object-oriented language but do have structs, which is an alternative to classes.

Executables are different for OS.

Widely used in Cloud infrastructure, dropbox, etc.

-----


### Setup

After installation, create a folder with name `HelloWorld` and create a `main.go` file in it. 

Now change your directory to the `HelloWorld` folder and run the following command in the bash
```bash
go mod init <name>
```

What does this file do?


In the `main.go`
```go
package main

import "fmt"

func main() {
	fmt.Println("Bye World")
}

```
To execute the file run `go run main.go` in the terminal.

----


### Types in Go lang.

Types are case sensitive, valirable type should be known in advance. Almost everything in Go lang is type.

+ String
+ Bool
+ Integer - uint8, uint64, int8, int64, uintptr
+ Floating - float32, float63
+ Complex
+ Arrays
+ Slices
+ Maps
+ Structs
+ Pointers
+ Others - Functions, Channels


----

### Builds in Go
Type `go env` in the command line and you will find a list of declared variables. One of them is `GOOS`.  It is different for different os
```go
// For Linux
GOOS = "linux"

// For Mac
GOOS = "darwin"

// For Windows
GOOS = "windows"
```

To build a file for a specific OS. Run the following command in the command line.
```bash
GOOS = "linux" go build
```

This build build the `main.go` file in the root directory.

--- 

### Memory in Golang.
In Goolang, memory allocation and deallocation happens automatically.

+ `new()`
	+ Allocate memory but no INIT
	+ Returns a memory address
	+ Zeroed usage - initiallyy we can not put any data.
+ `make()`
	+ Allocate memory and INIT
	+ One will get a memory address
	+ Non zeroed usage

----

### Pointers in Go.
Why does it exist?

Most of the language have some common problem, and that is that when we create a variable in a program which is nothing but a reference to a address. So when we use it, sometimes it creates a copy of the address and create irregularities in the program.

So using pointer, when we pass a resource as a pointer. It ensures that the actual value is passed from the memory.