package app

import (
	"banking/service"
	"encoding/json"
	// "encoding/xml"
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type CustomerHandler struct{
	service service.CustomerService
}


func (ch CustomerHandler) getAllCustomers(w http.ResponseWriter,r *http.Request){
	status:=r.URL.Query().Get("status")
	
	customers,err:=ch.service.GetAllCustomer(status)
	if err!=nil{
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		writeResponse(w,http.StatusOK,customers)
	}
}

func (ch CustomerHandler) getCustomers(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id:=vars["customer_id"]

	customer,err:=ch.service.GetCustomer(id)
	if err!=nil{
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		writeResponse(w,http.StatusOK,customer)
	}
}


func writeResponse(w http.ResponseWriter,code int,data interface{}){
	w.Header().Add("Content-Type","application/json")
		w.WriteHeader(code)
		if err:=json.NewEncoder(w).Encode(data);err!=nil{
			panic(err)
		}
}
