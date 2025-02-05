package domain

import "github.com/puttarajkoliwad/go_projects/banking_app/errs"

type Customer struct {
	Id string		`db:"customer_id"`
	Name string
	City string
	Zipcode string
	Dob string		`db:"date_of_birth"`
	Status string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, error)
	FindById(string) (*Customer, *errs.AppError)
}
