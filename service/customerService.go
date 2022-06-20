package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type CustomerService interface{
	GetAllCustomer(string)([]dto.CustomerResponse,*errs.AppError)
	GetCustomer(string)(*dto.CustomerResponse,*errs.AppError)

}

type DefaultCustomerService struct{
	repo domain.CustomerRepositoryDb
}

func (s DefaultCustomerService) GetAllCustomer(status string)([]dto.CustomerResponse,*errs.AppError){
	if status=="active"{
		status="1"
	}else if status=="inactive" {
		status="0"
	}else{
		status=""
	}
	var r []dto.CustomerResponse
	c,err:= s.repo.FindAll(status)
	if err!=nil{
		return nil,err
	}
	for _,v:=range c{
		r=append(r, v.ToDto())
	} 
	return r,nil

}

func (s DefaultCustomerService) GetCustomer(id string)(*dto.CustomerResponse,*errs.AppError){
	c,err:= s.repo.ById(id)
	if err!=nil{
		return nil,err
	}
	response:=c.ToDto()
	return &response,nil
}

func NewCustomerService(repository domain.CustomerRepositoryDb)DefaultCustomerService{
	return DefaultCustomerService{repository}
}