package accounts

import (
	"eventzourting/pkg/eventsourcing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {

	acc := new(Account)
	acc.changes = []eventsourcing.Event{}
	acc.Handle(&CreateCommand{InitialBalance: 100})
	acc.Handle(&DepositCommand{Amount: 100})
	acc.Handle(&DepositCommand{Amount: 50})
	acc.Handle(&WithdrawCommand{
			Amount: 20,
		})
	assert.Equal(t, 230, acc.Balance)
	assert.Equal(t, 4,len(acc.GetChanges()))

}
