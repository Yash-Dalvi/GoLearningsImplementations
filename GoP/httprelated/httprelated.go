package httprelated

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structs in Go

// When to use type in Go ?
// When declaring a variable without specifying an explicit type
// (either by using the := syntax or var = expression syntax), the variable's type is
// inferred from the value on the right hand side.

type Customers struct {
	Name       string
	Id         int
	Occupation string
}

func Starthttpthings() {
	// 2. Now I want to write a function which would take in desired path which will be mentioned as request
	//for the above server and it would written response based on a function that I define

	// Note - This end point needs to be written before the Listen and serve function is written!
	http.HandleFunc("/first", firstendpointfunction)
	http.HandleFunc("/customers", GetAllCustomers)

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
}

func firstendpointfunction(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello I am serving firstendpoint return function")

}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	/* Below is how one can initialize a struct and define the values of the variables in the struct as well.
	ross := Employee {
		firstName: "ross",
		lastName:  "Bing",
		fullTime:  true,
	}*/

	// I will declare a array of type Customers structs and initialize them

	CustomersArray := []Customers{
		{Name: "Customer1", Id: 1, Occupation: "IT Engineer"},
		{Name: "Customer2", Id: 2, Occupation: "MechTech Repar"},
	}

	// Below is how we set the ResponseWriter header.
	w.Header().Add("Content-Type", "application/json")

	// Below is how we convert the struct into a json response using the json encoder
	json.NewEncoder(w).Encode(CustomersArray)

}
