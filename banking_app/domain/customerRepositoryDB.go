package domain

import (
	"net/http"
	"database/sql"
	"time"
	"github.com/puttarajkoliwad/go_projects/banking_app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/puttarajkoliwad/go_projects/banking_app/errs"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, error) {
	var rows *sql.Rows
	var err error

	findAllSql := "select * from customers"

	if status != "" {
		findAllSql += " where status = ?"
		rows, err = d.client.Query(findAllSql, status)
	}else{
		rows, err = d.client.Query(findAllSql)
	}

	logger.Debug(findAllSql)

	if err != nil {
		logger.Error("Error fetching customers!" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dob, &c.Status)

		if err != nil {
			logger.Error("Error scanning customer rows" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (cr CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	q := "select * from customers where customer_id = ?"

	row := cr.client.QueryRow(q, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dob, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer does not exist!   " + err.Error())
			return nil, &errs.AppError{http.StatusNotFound, "Customer does not exist!"}
		}

		logger.Error("Error scanning customer details		" + err.Error())
		return nil, &errs.AppError{http.StatusInternalServerError, "Unexpected database error!"}
	}

	return &c, nil
}

func NewCustomerRepositoryDB() (*CustomerRepositoryDB) {
	client, err := sql.Open("mysql", "root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDB{
		client,
	}
}

