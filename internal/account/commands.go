package accounts

import "eventzourting/pkg/eventsourcing"

type CreateCommand struct {
	eventsourcing.WithGuid
	InitialBalance int
}
type WithdrawCommand struct {
	eventsourcing.WithGuid
	Amount int
	Description	string
}
type DepositCommand struct {
	eventsourcing.WithGuid
	Amount int
}
type CloseCommand struct {
	eventsourcing.WithGuid
}