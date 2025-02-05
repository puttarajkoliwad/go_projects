package domain

import "github.com/puttarajkoliwad/go_projects/banking_app/errs"
import "github.com/puttarajkoliwad/go_projects/banking_app/dto"


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

func(c Customer) statusAsText() string {
	status := "active"
	if c.Status == "0" {
		status = "inactive"
	}

	return status
}

func(c Customer) ToDto() (*dto.CustomerResponse) {
	return &dto.CustomerResponse {
		c.Id,
		c.Name,
		c.City,
		c.Zipcode,
		c.Dob,
		c.statusAsText(),
	}
}