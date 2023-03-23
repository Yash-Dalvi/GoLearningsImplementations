package service

import "GOP/hexagonal_udemy/domain"

// Now this package deals with the left hand side of the hexagonal,
// Every port is an interface in this hexagonal pattern, so
// we create a new port , service, with intention of an external adapter driving this
// service via connecting with this port!

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	// this is the utomost important central part of the hexagon, which has dependency
	// on the CustomerRepository port
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
