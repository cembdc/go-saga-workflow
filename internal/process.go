package order

import (
	"errors"
	wf "go-saga-workflow/pkg"
	"log"
)

type Order struct {
	Id       uint
	Product  string
	Quantity int
	Total    float64
}

func ProcessOrder(workflow *wf.Workflow) {

	order := &Order{
		Id:       123,
		Product:  "Laptop",
		Quantity: 1,
		Total:    1000.00,
	}

	workflow.AddStep("make_order", wf.SagaStep{Transaction: order.MakeOrder, Compensate: order.CancelOrder})
	workflow.AddStep("process_payment", wf.SagaStep{Transaction: order.ProcessPayment, Compensate: order.CancelPayment})
}

func (o *Order) MakeOrder() error {

	log.Printf("Order created: %v\n", o)

	return nil
}

func (o *Order) CancelOrder() error {
	log.Printf("Order cancelled: %d\n", o.Id)
	return nil
}

func (o *Order) ProcessPayment() error {
	return errors.New("payment failed")
}

func (o *Order) CancelPayment() error {
	log.Printf("Payment cancelled: %v\n", o.Total)
	return nil
}
