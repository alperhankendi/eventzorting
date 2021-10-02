package eventsourcing

type Command interface {
	MyId
}

type Event interface {
	MyId
}

type Aggregate interface {
	MyId
	When(Event)
	Apply(Event) []Event
	VersionUp()
}
