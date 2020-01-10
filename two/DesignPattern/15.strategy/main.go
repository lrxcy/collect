package main

import "fmt"

type PaymentContext struct {
	Name, CardID string
	Money        int
	payment      PaymentStrategy
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

func NewPaymentContext(name, cardid string, money int, payment PaymentStrategy) *PaymentContext {
	return &PaymentContext{
		Name:    name,
		CardID:  cardid,
		Money:   money,
		payment: payment,
	}
}

func (p *PaymentContext) Pay() {
	p.payment.Pay(p)
}

type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s\n", ctx.Money, ctx.Name, ctx.CardID)
}

func main() {
	ctxc := NewPaymentContext("Ada", "", 123, &Cash{})
	ctxc.Pay()

	ctxb := NewPaymentContext("Bob", "002", 888, &Bank{})
	ctxb.Pay()
}
