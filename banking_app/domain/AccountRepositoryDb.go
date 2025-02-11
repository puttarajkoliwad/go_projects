package domain

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/puttarajkoliwad/go_projects/banking_app/errs"
	"github.com/puttarajkoliwad/go_projects/banking_app/logger"
	"net/http"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func(ar AccountRepositoryDb) Save(acc Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := ar.client.Exec(sqlInsert, acc.CustomerId, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status)
	if err != nil {
		logger.Error("Error inserting account into DB " + err.Error())
		return nil, &errs.AppError{http.StatusInternalServerError, "Account creation failed!"}
	}

	id, err := result.LastInsertId()
	if err != nil {
		errMsg := "Unable retrive created account id"
		logger.Error(errMsg + " " + err.Error())
		return &acc, nil
	}

	acc.Id = strconv.FormatInt(id, 10)
	return &acc, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) (*AccountRepositoryDb) {
	return &AccountRepositoryDb{
		client,
	}
}