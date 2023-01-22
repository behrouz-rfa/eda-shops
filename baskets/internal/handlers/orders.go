package handlers

import (
	"eda-shops/baskets/internal/application"
	"eda-shops/baskets/internal/domain"
	"eda-shops/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
