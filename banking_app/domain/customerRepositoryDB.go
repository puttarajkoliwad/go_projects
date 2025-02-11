package domain

import (
	"net/http"
	"database/sql"
	"github.com/jmoiron/sqlx"
	// "time"
	"github.com/puttarajkoliwad/go_projects/banking_app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/puttarajkoliwad/go_projects/banking_app/errs"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, error) {
	var err error

	customers := make([]Customer, 0)
	findAllSql := "select * from customers"

	if status != "" {
		findAllSql += " where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}else{
		err = d.client.Select(&customers, findAllSql)
	}

	logger.Debug(findAllSql)

	if err != nil {
		logger.Error("Error fetching customers!" + err.Error())
		return nil, err
	}

	// customers := make([]Customer, 0)
	// if err := sqlx.StructScan(rows, &customers); err != nil {
	// 	logger.Error("Error scannign rows		" + err.Error())
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dob, &c.Status)

	// 	if err != nil {
	// 		logger.Error("Error scanning customer rows" + err.Error())
	// 		return nil, err
	// 	}
	// 	customers = append(customers, c)
	// }

	return customers, nil
}

func (cr CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	q := "select * from customers where customer_id = ?"

	// row := cr.client.QueryRow(q, id)

	var c Customer
	err := cr.client.Get(&c, q, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dob, &c.Status)

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

func NewCustomerRepositoryDB(client *sqlx.DB) (*CustomerRepositoryDB) {
	return &CustomerRepositoryDB{
		client,
	}
}

