// Packages in Go - What would they be?!
// They might be collection of module, a module typically would contain code written by someone else,
// the functions in the modules and someone else would be using this package, so that this someone would save
// the effort of writing the same code written by someone else!

// The concept of go mod init - this go.mod is dependency management file,
// in this file, the track of all the dependent packages are kept!
package main

import (
	"fmt"
	"net/http"
	"test"
)

func main() {
	fmt.Print("HEllo world from Yash.")

	// 2. Now I want to write a function which would take in desired path which will be mentioned as request
	//for the above server and it would written response based on a function that I define

	// Note - This end point needs to be written before the Listen and serve function is written!
	http.HandleFunc("/first", firstendpointfunction)

	// 1. I want to create a simple http server, listening on a port and serving a simple response on /firstendpoint.
	// They way to do this is to use inbuilt http go package.

	http.ListenAndServe("localhost:7777", nil)
	// the above function does the required, the second argument it takes is nil, which tells the function
	// to use default server multiplexer.

	// A server multiplexer is like a big routing box, which has info of all routes registered and it
	// processess the requests based on its table!

	// We usually avoid using the default sever mux because it has very limited functions and we will
	// need to write very cumbersome code to use the default server mux.
	// Like if any variable is passed in the url, then default server mux does not provide us the funcitonality
	// to extract the variable out of the url.

	// 3. Working with various packages

	test.Simple()

}

func firstendpointfunction(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello I am serving firstendpoint return function")

}
