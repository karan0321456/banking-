package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"


	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string)([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	var err error
	if status==""{
		findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err=d.client.Select(&customers,findAllSql)
	}else{
		findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where status=?"
		err=d.client.Select(&customers,findAllSql,status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil,errs.NewUnexpectedError("Unexpected Database error" )
	}
	return customers, nil
}

func (d CustomerRepositoryDb)ById(id string)(*Customer,*errs.AppError){
	var c Customer
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?"
	err:=d.client.Get(&c,customerSql,id)
	if err!=nil{
		if err == sql.ErrNoRows{
			return nil,errs.NewNotFoundError("Customer Not Found")
		}else{
			
			logger.Error("Error while scanning customer table " + err.Error())
			return nil,errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c,nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB)CustomerRepositoryDb{
	
	return CustomerRepositoryDb{client: dbClient}

}