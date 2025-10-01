package main

import (
	"fmt"
	"math/rand"
)

type PaymentMethod interface {
	makePayment(amount int) (bool, error)
}

type CreditCard struct {
	no         string
	cvv        int
	expiryDate string
}

func (cc *CreditCard) makePayment(amount int) (bool, error) {
	// TODO: making payment
	fmt.Printf("Made payment amount %v with credit card %v \n", amount, cc.no)
	return true, nil
}

type UPI struct {
	upiId  string
	upiPin int
}

func (u *UPI) makePayment(amount int) (bool, error) {
	// TODO: make payment
	fmt.Printf("made payment amount %v with upi_id %v \n", amount, u.upiId)
	return true, nil
}

type User struct {
	user_id       string
	paymentMethod PaymentMethod
}

func (u *User) doShopping() {
	// TODO: Select items
	// TODO: Ask Prics
	amount := rand.Int() % 10000
	// make Payment
	_, err := u.paymentMethod.makePayment(amount)
	if err != nil {
		fmt.Printf("Error in making payment %s", err.Error())
	}
}

func main() {
	var paymentMethod PaymentMethod
	paymentMethod = &UPI{
		upiId:  "1234",
		upiPin: 0,
	}
	user := User{
		user_id:       "1234",
		paymentMethod: paymentMethod,
	}
	paymentMethod = &CreditCard{
		no:         "124-134-679",
		cvv:        123,
		expiryDate: "12445774",
	}
	user.doShopping()
	// you switch the payment method from UPI to CreditCard — demonstrating runtime polymorphism.
	user.paymentMethod = paymentMethod
	user.doShopping()
}
