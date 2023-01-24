package domain

const (
	StoreCreatedEvent               = "stores.StoreCreated"
	StoreParticipationEnabledEvent  = "stores.StoreParticipationEnabled"
	StoreParticipationDisabledEvent = "stores.StoreParticipationDisabled"
	StoreRebrandedEvent             = "stores.StoreRebranded"
)

type StoreCreated struct {
	Name     string
	Location string
}

// Key implements registry.Registerable
func (StoreCreated) Key() string { return StoreCreatedEvent }

func (StoreCreated) EventName() string { return "stores.StoreCreated" }

type StoreParticipationToggled struct {
	Participating bool
}
type StoreRebranded struct {
	Name string
}

// Key implements registry.Registerable
func (StoreRebranded) Key() string { return StoreRebrandedEvent }
