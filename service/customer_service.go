package service

import (
	"Hexagonal/repository"
	"database/sql"
	"errors"
	"log"
)

type customerService struct {
	//เรียกตัวปลักที่สร้างไว้
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) customerService {
	return customerService{customerRepository: customerRepository}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	CustResponses := []CustomerResponse{}

	for _, customer := range customers {
		CustResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		CustResponses = append(CustResponses, CustResponse)
	}

	return CustResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}

	CustResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &CustResponse, nil
}
