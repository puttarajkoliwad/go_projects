package domain

import "github.com/puttarajkoliwad/go_projects/banking_app/errs"

type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	Dob string
	Status string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, *errs.AppError)
}
