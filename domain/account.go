package domain

import (
	"banking/dto"
	"banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func(a Account) ToNewAccountResponseDto() dto.NewAccountResponse{
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account)(*Account,*errs.AppError)
	SaveTransaction(transaction Transaction)(*Transaction,*errs.AppError)
	FindBy(accountId string)(*Account,*errs.AppError)
}

func (a Account) CanWithdraw(amount float64)bool{
	if a.Amount < amount {
		return false
	}
	return true
}

func NewAccount(customerID,accountType string,amount float64)Account{
	return Account{
		CustomerId: customerID,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount: amount,
		Status: "1",
	}
}

