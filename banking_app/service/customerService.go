package service

import (
	"github.com/puttarajkoliwad/go_projects/banking_app/domain"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) (CustomerService) {
	return DefaultCustomerService{repository}
}