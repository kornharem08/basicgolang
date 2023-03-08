package service

import (
	"Hexagonal/errs"
	"Hexagonal/logs"
	"Hexagonal/repository"
	"database/sql"
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
		logs.Error(err)
		return nil, errs.NewUnexpectedError("unexpected error")
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
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError("unexpected error")
	}

	CustResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &CustResponse, nil
}
