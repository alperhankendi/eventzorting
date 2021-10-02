package accounts

import (
	"eventzourting/pkg/eventsourcing"
	"fmt"
)
const (
	Open = State("Open")
	Closed = State("Closed")
)
type (
	Account struct {
		Id 		eventsourcing.WithGuid
		Balance int
		State State
		Version int
		changes []eventsourcing.Event
	}
	State string
)
func (a *Account) When(e eventsourcing.Event) {

		switch event:=e.(type) {
		case *AccountCreated:
			a.Balance = event.InitialBalance
			a.State = Open
			break
		case *Deposited: a.Balance +=event.Amount
			break
		case *Withdrew: a.Balance -= event.Amount
			break
		case *AccountClosed: a.State = Closed
			break
		case *WithdrawFailedEvent:
			break
		default:
			panic(fmt.Sprintf("unreconized event. %v",event))
		}
}
func (a *Account) Handle(c eventsourcing.Command){

	var event eventsourcing.Event
	switch cmd := c.(type) {
		case *CreateCommand:
			event = &AccountCreated{InitialBalance: cmd.InitialBalance}
			break
		case *DepositCommand:
			event = &Deposited{Amount:   cmd.Amount,}
			break
		case *WithdrawCommand:
			if a.Balance < cmd.Amount{
				event = &WithdrawFailedEvent{
					Message: fmt.Sprintf("Failed to withdraw operation. Balance:%d, Requested:%d",a.Balance,cmd.Amount),
				}
			}else {
				event = &Withdrew{
					Amount:      cmd.Amount,
					Description: cmd.Description,
				}
			}
			break
		case *CloseCommand:
			event = &AccountClosed{}
			break
	}
	event.SetGuid(c.GetGuid())
	a.Apply(event)
}
//Don't like 'generic "Handle" function', alternative way...
//func (a *Account) Deposit(cmd DepositCommand)  {
//	var event eventsourcing.Event
//	event = &Deposited{Amount:   cmd.Amount,}
//	event.SetGuid(cmd.GetGuid())
//	a.Apply(event)
//}

func (a *Account) Apply(event eventsourcing.Event)  {
	a.When(event)
	//uncommited changes... will be saved eventstore..
	//reminder0 : no pump aggregate version, optimistic lock rule
	//reminder1 : check aggregate state before save the event! aggregate state must be valid.
	a.changes = append(a.changes,event)
}

func (a *Account) GetChanges() []eventsourcing.Event{
	return a.changes
}
func (a *Account) ClearChanges() {
	a.changes= []eventsourcing.Event{}
}
func (a *Account) ShowChanges() {

	for _,e :=range a.changes{
		fmt.Printf("Event:%T}\t : %v\n",e, e)
	}
}

func Load(guid eventsourcing.Guid, events []eventsourcing.Event) *Account {
	a:= NewAccount()

	for _,e :=range events{
		a.When(e)
		a.Version++
	}
	a.Id.SetGuid(guid)
	return a
}

func NewAccount() *Account {
	return &Account{
		changes: []eventsourcing.Event{},
		Id: eventsourcing.WithGuid{Guid: eventsourcing.NewGuid()},
	}
}