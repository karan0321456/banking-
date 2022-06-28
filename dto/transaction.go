package dto

import "banking/errs"


const WITHDRAWL ="withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct{
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId   	string  `json:"customer_id"`
}

func (r TransactionRequest)IstransactionTypeWithdrawal()bool{
	return r.TransactionType == WITHDRAWL
}

func (r TransactionRequest)IstransactionTypeDeposit()bool{
	return r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError{
	if !r.IstransactionTypeWithdrawal() && !r.IstransactionTypeDeposit(){
	return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}else{
		if r.Amount <0{
			return errs.NewValidationError("Amount cannot be less than zero")
		}
	}
	return nil
}


type TransactionResponse struct{
	TransactionId  	string	`json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}