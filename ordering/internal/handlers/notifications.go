package handlers

import (
	"eda-shops/internal/ddd"
	"eda-shops/ordering/internal/application"
	"eda-shops/ordering/internal/domain"
)

func RegisterNotificationHandlers(notificationHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderCreated{}, notificationHandlers.OnOrderCreated)
	domainSubscriber.Subscribe(domain.OrderReadied{}, notificationHandlers.OnOrderReadied)
	domainSubscriber.Subscribe(domain.OrderCanceled{}, notificationHandlers.OnOrderCanceled)
}
