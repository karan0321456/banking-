package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(){
	r:= mux.NewRouter()

	//wiring
	// ch:=CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub()) }
	ch:=CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb()) }

	//routes
	r.HandleFunc("/customers",ch.getAllCustomers).Methods("GET")
	r.HandleFunc("/customers/{customer_id:[0-9]+}",ch.getCustomers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8457",r))
}
