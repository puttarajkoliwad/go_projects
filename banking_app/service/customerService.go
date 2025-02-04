package service

import (
	"github.com/puttarajkoliwad/go_projects/banking_app/domain"
	"github.com/puttarajkoliwad/go_projects/banking_app/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, error) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindById(id)
}

func NewCustomerService(repository domain.CustomerRepository) (CustomerService) {
	return DefaultCustomerService{repository}
}