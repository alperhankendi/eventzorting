package accounts

import "eventzourting/pkg/eventsourcing"

type AccountCreated struct {
	eventsourcing.WithGuid
	InitialBalance int
}
type Withdrew struct {
	eventsourcing.WithGuid
	Amount int
	Description string
}
type Deposited struct {
	eventsourcing.WithGuid
	Amount int
}
type AccountClosed struct {
	eventsourcing.WithGuid
}

type WithdrawFailedEvent struct {
	eventsourcing.WithGuid
	Message string
}
