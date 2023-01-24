package handlers

import (
	"eda-shops/internal/ddd"
	"eda-shops/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, invoiceHandlers)
}
