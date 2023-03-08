package service

import (
	"Hexagonal/errs"
	"Hexagonal/logs"
	"Hexagonal/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) accountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {

	if request.Amount < 5000 {
		return nil, errs.NewValidationError("Amount at least 5,000")
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("unexpected error")
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}

func (s accountService) GetAccounts(customerId int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError("unexpected error")
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return responses, nil
}
