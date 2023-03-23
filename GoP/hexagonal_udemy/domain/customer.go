package domain

// According to the hexagonal architecture, below struct is kind of a domain object,
// the core object of an application one would say.

type Customer struct {
	Name        string
	Id          string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// Below is defination of one such port, implemented at the boundary on the server side
// of hexagonal architecture.
// According to the concept of port, any adapter which implements the functions defined in
// the port should be properly able to use/connect to this port!
type CustomerRepository interface {

	// This below function is expected to be implemented by all the adapters
	// which want to connect to this port!
	FindAll() ([]Customer, error)
}

// Below is now creating a Customer Adapter named CustomerRepositoryStub

/*type CustomerRepositoryStub struct {
	customers []Customer
	// Now this adapter should implement FindAll interface, which is done below
}*/
// The same actual implementation of above is found in file customerrepositorystub
