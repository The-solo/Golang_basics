package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line


func updateBalance(c *customer,t transaction) (error){
	//cheak if the pointer is null.
	if c == nil {
		return errors.New("Null pointer")
	}
	switch t.transactionType {
	case "deposit":
		(c.balance) += t.amount //automatic dereferencing.
		return nil
	case "withdrawal":
		if (c.balance) >= t.amount{
			(c.balance) -= t.amount
			return nil
		} else {
			return errors.New("insufficient funds")
		}
	default:
		return errors.New("unknown transaction type")
	}	
}
