package main

// Packages in Go - What would they be?!
// They might be collection of module, a module typically would contain code written by someone else,
// the functions in the modules and someone else would be using this package, so that this someone would save
// the effort of writing the same code written by someone else!

// The concept of go mod init - this go.mod is dependency management file,
// in this file, the track of all the dependent packages are kept!

import (
	//"GOP/hello"
	//"GOP/httprelated"
	"fmt"
	//"GOP/muxpack"
	"GOP/interfaceshexagonal"
)

func main() {

	// I can below package calls if they are not needed for now!!!
	fmt.Println("Hello from main")
	//httprelated.Starthttpthings()
	//hello.Testing()
	//muxpack.StartMux()
	//interfaceshexagonal.StartInterface()
	interfaceshexagonal.StartInterface2()

	// Need to continue understanding hexagonal architecture maybe somepoint later.
	// https://ibm-learning.udemy.com/course/rest-based-microservices-api-development-in-go-lang/learn/lecture/22579888#overview
	// above is the url to do so.

}
