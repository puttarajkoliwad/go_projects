package domain

import "github.com/puttarajkoliwad/go_projects/banking_app/errs"

type Account struct {
	Id string
	CustomerId string
	OpeningDate string
	AccountType string
	Amount float64
	Status string
}

type AccountRespository interface {
	Save(Account) (*Account, *errs.AppError)
}