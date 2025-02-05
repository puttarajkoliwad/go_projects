package service

import (
	"github.com/puttarajkoliwad/go_projects/banking_app/domain"
	"github.com/puttarajkoliwad/go_projects/banking_app/dto"
	"github.com/puttarajkoliwad/go_projects/banking_app/errs"
)
import lop "github.com/samber/lo/parallel"

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, error)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	customersDto := lop.Map(customers, func(c domain.Customer, _ int) dto.CustomerResponse {
		return *c.ToDto()
	})

	return customersDto, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToDto()

	return resp, nil
}

func NewCustomerService(repository domain.CustomerRepository) (CustomerService) {
	return DefaultCustomerService{repository}
}