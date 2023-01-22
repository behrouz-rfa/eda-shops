package handlers

import (
	"eda-shops/depot/internal/application"
	"eda-shops/depot/internal/domain"
	"eda-shops/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompleted{}, orderHandlers.OnShoppingListCompleted)
}
