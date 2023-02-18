## Go lang

*Go lang* is a compiled language, but it can run files directly, without VM. Similarity with lots of languages like C, Java, and Pascal.

It is not an object-oriented language but does have structures, which is an alternative to classes.

Executables are different for OS.

Widely used in Cloud infrastructure, dropbox, etc.

-----


## Setup

After installation, create a folder with the name `HelloWorld` and create a `main.go` file in it.

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


## Types in Go lang.

Types are case-sensitive, and variable types should be known in advance. Almost everything in Go lang is typed.

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

Example :
```go
	var smallfFloat float64 = 255.0131564651315
	fmt.Println(smallfFloat)
	fmt.Printf("Variable is of type: %T \n", smallfFloat)
```

----

## Builds in Go
Type `go env` in the command line and you will find a list of declared variables. One of them is `GOOS`. It is different for different os
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

This build builds the `main.go` file in the root directory.

---

## User Input
`bufio` and `os` is the package that helps getting the user input and output.

```go
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")

	// Comma Ok or Error Ok Synatx
	input, _ := reader.ReadString('\n')
```

## Type Conversion

**Packages**
```go
import (
	"strconv"
	"strings"
)

```

```go
	age, _ := strconv.ParseInt(strings.TrimSpace(input_age), 10, 64)
	salary, err := strconv.ParseFloat(strings.TrimSpace(input_salary), 64)
```

------

## Time Package
```go
import (
	"fmt"
	"time"
)
```

```go
presentTime := time.Now()
	fmt.Print(presentTime)
```

This outputs - 2023-02-06 22:05:12.8190491 +0530 IST m=+0.001639001

We need to format the time in Month-Date-Year Hour:Minute:Second Day

--------

## Memory in Golang.
In Goolang, memory allocation, and deallocation happen automatically.

+ `new()`
    + Allocate memory but no INIT
    + Returns a memory address
    + Zeroed usage - initially we can not put any data.

+ `make()`
    + Allocate memory and INIT
    + One will get a memory address
    + Non-zeroed usage


----

## Slices 

```go
// Declaration
var avengersList = []string{"Iron Main", "Hulk", "Winter Soldier"}

// Another Method
highScores := make([]int, 4)
highScores[0] = 234

// Append
highScores = append(highScores, 123, 125, 124, 128, 256)

// Indexing
var index int = 2
courses = append(courses[:index], courses[index+1:]...)

```

-----

## Maps

```go
// Declaration
language := make(map[string]string)

// Loop on a Map
for key, value := range language {
	fmt.Printf("For key %v, value is: %v\n", key, value)
}

// Access a value using key
language["JS"]

// Delete a value corresponnding to a key.
delete(language, "RB")
```

----

## Structs
```go
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


// Create an instance
nikhil := User{"Nikhil", "nikhil25803@gmail.com", true, 18}


// Accessing
fmt.Printf("Name is: %v and Email is: %v", nikhil.Name, nikhil.Email)
```

-----



## Pointers in Go.
Why does it exist?

Most languages have some common problem, and that is that when we create a variable in a program that is nothing but a reference to an address. So when we use it, sometimes it creates a copy of the address and creates irregularities in the program.

So using a pointer, when we pass a resource as a pointer. It ensures that the actual value is passed from the memory.

```go
	myNumber := 25
	var ptr2 = &myNumber
	fmt.Println("value of the pointer with value is: ", ptr2)
```

-------

## Random Number

```go
import (
	"fmt"
	"math/rand"
	"time"
)

rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)
```

------


## Switch Case

```go
switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
    ...
	default:
		fmt.Println("What was that bro!")
	}
```

-----

## Functions

```go
// Create a function
func adder(valOne int, val2 int) int {
	return valOne + val2
}

// Call the function
result := adder(3, 5)
fmt.Println("Result is: ", result)
```

OR

```go
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}


proRes := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
fmt.Println("Sum of the slice is: ", proRes)
```

-----------

## Defer in Golang.
In `Golang`, the `defer` keyword is used to *delay* the execution of a function or a statement until the nearby function returns.

In simple words, `defer` will move the execution of the statement to the very end of a function.

```go
package main

import "fmt"

func main() {
fmt.Println("Welcome to Defers in Golang")
greet()
}

func greet() {
defer fmt.Println("World")
fmt.Println("Hello")
}

```
When we call the `greet()` function, the `fmt.Println("World")` should have been invoked first but as it contains the keyword `defer`, it will delay it and execute `fmt.Println("Hello")` first.

> Output
```bash
Welcome to Defers in Golang
Hello
World
```

------

## Error check

```go
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
```

----


## Write into a file
```go
content := "This is a test content to be written in a txt file"

	file, err := os.Create("./myGofile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
```

