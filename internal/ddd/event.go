package ddd

import (
	"github.com/google/uuid"
	"time"
)

type (
	EventPayload interface{}
	Event        interface {
		IDer
		EventName() string
		Payload() EventPayload
		MetaData() Metadata
		OccurredAt() time.Time
	}
	event struct {
		Entity
		payload    EventPayload
		metadata   Metadata
		occurredAt time.Time
	}
)

func NewEvent(name string, payload EventPayload, options ...EventOption) event {
	return newEvent(name, payload, options...)
}
func newEvent(name string, payload EventPayload, options ...EventOption) event {
	evt := event{
		Entity:     NewEntity(uuid.New().String(), name),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}
	for _, option := range options {
		option.configureEvent(&evt)
	}
	return evt
}

var _ Event = (*event)(nil)

func (e event) EventName() string {
	//TODO implement me
	panic("implement me")
}

func (e event) Payload() EventPayload {
	//TODO implement me
	panic("implement me")
}

func (e event) MetaData() Metadata {
	//TODO implement me
	panic("implement me")
}

func (e event) OccurredAt() time.Time {
	//TODO implement me
	panic("implement me")
}
