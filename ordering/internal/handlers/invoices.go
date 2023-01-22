package handlers

import (
	"eda-shops/internal/ddd"
	"eda-shops/ordering/internal/application"
	"eda-shops/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
