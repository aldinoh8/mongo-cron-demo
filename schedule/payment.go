package schedule

import (
	"example/model"
	"example/repository"
	"fmt"
)

type PaymentHandler struct {
	Repository repository.Payment
}

func NewPayment(r repository.Payment) PaymentHandler {
	return PaymentHandler{Repository: r}
}

func (handler PaymentHandler) UpdateInvoiceStatus() {
	fmt.Println("Triggered update invoice status")
	invoices, err := handler.Repository.FindNotPaid()
	if err != nil {
		fmt.Println("failed to get not paid invoices")
		return
	}

	for _, inv := range invoices {
		handler.UpdateStatus(inv)
	}
}

func (handler PaymentHandler) UpdateStatus(invoice model.Invoice) {
	payment, err := handler.Repository.PaymentFindByInvoiceId(invoice.ID)
	if err != nil {
		fmt.Printf("failed to update %v; error: %s\n", invoice.ID, err.Error())
		return
	}

	err = handler.Repository.PaidInvoice(invoice.ID)
	if err != nil {
		fmt.Printf("failed to update %v; error: %s\n", invoice.ID, err.Error())
		return
	}

	fmt.Printf("success update status to PAID INVOICE ID:%v | PAYMENT ID: %v\n", invoice.ID, payment.ID)
}
