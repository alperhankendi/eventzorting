package eventsourcing

import "github.com/google/uuid"

type MyId interface {
	GetGuid() Guid
	SetGuid(Guid)
}
type WithGuid struct {
	Guid Guid
}
func (e *WithGuid) SetGuid(g Guid) {
	e.Guid = g
}
func (e *WithGuid) GetGuid() Guid {
	return e.Guid
}
type Guid string

func NewGuid() Guid {
	return Guid(uuid.New().String())
}

