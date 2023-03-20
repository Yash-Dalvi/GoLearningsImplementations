package hello

import "fmt"

// This is special headache requirement of Go
// Point to note is that in the below function name, "T" should be capital, only then it can be used outside the
// package!
func Testing() {
	fmt.Print("This is coming from testing")
}
