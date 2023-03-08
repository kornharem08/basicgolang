package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customer := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New York", ZipCode: "110075", Status: 1},
		{CustomerID: 1002, Name: "Fuca", City: "New York", ZipCode: "110075", Status: 1},
	}

	return customerRepositoryMock{customer}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
