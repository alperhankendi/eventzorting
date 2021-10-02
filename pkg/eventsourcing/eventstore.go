package eventsourcing

import (
	"errors"
	"fmt"
	"math"
)
type EventStore interface {
	Find(guid Guid) (events []Event, version int)

	Update(guid Guid, version int, events []Event) error

	GetEvents(offset int, batchSize int) []Event
}

type MemoryStore struct {
	store map[Guid][]Event
	events []Event
}
func (es *MemoryStore) Find(guid Guid) ([]Event, int) {
	events := es.store[guid]
	return events, len(events)
}

func (es *MemoryStore) Update(guid Guid, version int, events []Event) error{
	changes, ok := es.store[guid]
	if !ok {
		changes = []Event{}
	}
	if len(changes) == version {
		for _, event := range events {
			event.SetGuid(guid)
		}
		es.appendEvents(events)
		es.store[guid] = append(changes, events...)
	} else {
		return errors.New(
			fmt.Sprintf("WrongExpectedVersionException,entity has version %v, but event store %v", version, len(changes)))
	}
	return nil
}

func (es *MemoryStore) GetEvents(offset int, batchSize int) []Event {
	until := int(math.Min(float64(offset + batchSize), float64(len(es.events))))
	return es.events[offset:until]
}

func NewInMemoryStore() *MemoryStore {
	return &MemoryStore{
		store:map[Guid][]Event{},
		events:make([]Event, 0),
	}
}

func (es *MemoryStore) appendEvents(events []Event) {
	es.events = append(es.events, events...)
}