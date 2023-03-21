package main

// Packages in Go - What would they be?!
// They might be collection of module, a module typically would contain code written by someone else,
// the functions in the modules and someone else would be using this package, so that this someone would save
// the effort of writing the same code written by someone else!

// The concept of go mod init - this go.mod is dependency management file,
// in this file, the track of all the dependent packages are kept!

import (
	"GOP/hello"
	"GOP/httprelated"
	"fmt"
)

func main() {

	// I can below package calls if they are not needed for now!!!
	fmt.Print("Hello")
	httprelated.Starthttpthings()
	hello.Testing()

}
