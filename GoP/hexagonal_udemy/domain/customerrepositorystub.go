package domain

type CustomerRepositoryStub struct {
	customers []Customer
	// Now this adapter should implement FindAll interface, which is done below
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// Now we will just write a helper function that would initialize customers data for us

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "firstCustomer", City: "Agra", Zipcode: "4100", Status: "Active", DateOfBirth: "01/01/2000"},
		{Id: "2", Name: "secondCustomer", City: "NewDelhi", Zipcode: "4110", Status: "Active", DateOfBirth: "23/01/2000"},
	}
	return CustomerRepositoryStub{customers}
}

// Next to do is to understand the complete dependency of the diagram flow.
// Understand what is dependency injection!!
// For dependency injection I have created a separate Word Document!!

