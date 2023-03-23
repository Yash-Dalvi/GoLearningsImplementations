package interfaceshexagonal

import "fmt"

// Before understanding the hexagonal architecture , it is pre-requisite to understand
// and probably use interfaces in Go.
// Currently, as per my current understanding, interface is something that two different
// modules use to talk to each other.
// So now , at first, I will study the interfaces in detail then would visit the
// hexagonal architecture.

// I have written a beautiful doc explaining what exactly interface is and how it
// can be used , all thanks to a great well written blog on medium!!

// I am all set to write and implement a interface program in Go!!

type Duck interface {
	Quack()
	Walk()
}

type Donald struct {
	Name string
}

func (d Donald) Quack() {
	fmt.Println("I Quack")
}

func (d Donald) Walk() {
	fmt.Println("I walk")
}

func StartInterface() {
	d := new(Donald)
	d.Quack()
	d.Walk()
}

// Will now write a code which would actually demonstrate how interfaces can be useful in
// Go and how they help in scaliability and re-usiability and expansion as well!
// And how the concept of run time polymorphism is used in the case of interfaces!
