package main

import (
	"fmt"
	"math"
)

type PaymentMethod interface {
	ProcessPayment(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f using Credit Card ending in %s", amount, c.CardNumber[len(c.CardNumber)-4:])
}

type PayPal struct {
	Email string
}

func (p PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f using PayPal account %s", amount, p.Email)
}

type ShoppingCart struct {
	items []float64
}

func (sc *ShoppingCart) AddItem(price float64) {
	sc.items = append(sc.items, price)
}

func (sc ShoppingCart) TotalAmount() float64 {
	total := 0.0
	for _, price := range sc.items {
		total += price
	}
	return math.Round(total*100) / 100
}

func (sc ShoppingCart) Checkout(pm PaymentMethod) {
	total := sc.TotalAmount()
	fmt.Println(pm.ProcessPayment(total))
}

func main() {
	cart := ShoppingCart{}
	cart.AddItem(29.99)
	cart.AddItem(49.50)
	cart.AddItem(15.75)

	fmt.Printf("Total amount in cart: $%.2f\n", cart.TotalAmount())

	creditCard := CreditCard{CardNumber: "1234567812345678"}
	payPal := PayPal{Email: "user@example.com"}

	cart.Checkout(creditCard)
	cart.Checkout(payPal)
}
