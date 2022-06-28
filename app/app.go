package app

import (
	"banking/domain"
	"banking/service"
	"banking/logger"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func Start() {

	sanityCheck()
	r := mux.NewRouter()

	dbClient := getDbClient()

	customerRepositoryDb  := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	//wiring
	// ch:=CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub()) }
	ch := CustomerHandler{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	//routes
	r.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	r.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomers).Methods("GET")
	r.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods("Post")
	r.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods("Post")

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), r))
}


func getDbClient() *sqlx.DB{
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db:= os.Getenv("DB")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open(db, dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 10)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	logger.Info("Database is Connected")
	return client
}