package main

import "fmt" // Double quotes are important!
// Double quotes are important!

func main() {
	fmt.Println("Hi there!")
}

////////////////////////////////////////////////////////////////////////
//                               Go CLI                               //
////////////////////////////////////////////////////////////////////////

// - go build -> Compiles a bunch of go source files
// - go run -> compiles and executes one or two files
// - go fmt -> formats all the code in each file in the current directory
// - go install -> compiles and 'installs' a package
// - go get -> downloads the raw source code of someone else's package
// - go test -> runs any tests associated with the current project

////////////////////////////////////////////////////////////////////////
//                            Package Main                            //
////////////////////////////////////////////////////////////////////////

// What does a "package main" mean?

// package == project == workspace

// every file must declare the package it belongs to at the top

// there are two different types of packages
// - Executable
//     - Generates a file that we can run

// - Reusable
//     - Used as 'helpers'. Goo for reusable logic

// it is the name of the package that determines whether you're making an
// executable or dependency type package -> so package 'main' is sacred, it is
// always used as a executable package. Note that the main package MUST have a
// function called main w/in it

////////////////////////////////////////////////////////////////////////
//                         import statements                          //
////////////////////////////////////////////////////////////////////////

// - Just like importing any library / package......
// - fmt is a standard package like stdio?

////////////////////////////////////////////////////////////////////////
//                                func                                //
////////////////////////////////////////////////////////////////////////

// - just like functions anywhere else

////////////////////////////////////////////////////////////////////////
//                     How is main.go organized?                      //
////////////////////////////////////////////////////////////////////////

// 1) package delcaration `package main`

// 2) import other packages we need `import "fmt"`

// 3) Declare functions, tell go to do things `func main()......`
