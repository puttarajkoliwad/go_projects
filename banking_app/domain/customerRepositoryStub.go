package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1", "Rahul", "Blr", "560085", "2000-01-01", "1"},
		{"2", "Chetan", "Blr", "560085", "2000-01-01", "1"},
	}

	return CustomerRepositoryStub {
		customers,
	}
}
