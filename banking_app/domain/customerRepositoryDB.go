package domain

import (
	"database/sql"
	"time"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllSql := "select * from customers"

	rows, err := d.client.Query(findAllSql)	
	if err != nil {
		log.Println("Error fetching customers!" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dob, &c.Status)

		if err != nil {
			log.Println("Error scanning customer rows" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (cr CustomerRepositoryDB) FindById(id string) (*Customer, error) {
	sql := "select * from customers where customer_id = ?"

	row := cr.client.QueryRow(sql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dob, &c.Status)

	if err != nil {
		log.Println("Error scanning customer details", err.Error())
		return nil, err
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

