package main

import (
	accounts "eventzourting/internal/account"
	es "eventzourting/pkg/eventsourcing"
	"fmt"
)

func main(){

	var store = es.NewInMemoryStore()

	//create&update
	account1 := accounts.NewAccount()
	account1.Handle(&accounts.CreateCommand{
		WithGuid:       es.WithGuid{Guid: es.NewGuid()},
		InitialBalance: 10,
	})
	fmt.Printf("- Open account 1 with balance 10:\tOK\n")
	account1.Handle(&accounts.DepositCommand{
		WithGuid:       es.WithGuid{Guid: es.NewGuid()},
		Amount:   50,
	})
	account1.Handle(&accounts.DepositCommand{
		WithGuid:       es.WithGuid{Guid: es.NewGuid()},
		Amount:   40,
	})
	account1.Handle(&accounts.WithdrawCommand{
		WithGuid:    es.WithGuid{Guid: es.NewGuid()},
		Amount:      10,
		Description: "time is money, my friend",
	})
	fmt.Printf("\nAccount-> list uncommited changes:\n")
	account1.ShowChanges()
	err:=store.Update(account1.Id.GetGuid(), account1.Version, account1.GetChanges())
	account1.ClearChanges()
	if err != nil {
		fmt.Printf("Failed to update aggregate. \tError: %v",err)
	}
	fmt.Printf("Current Balance: %d\t State:%s\n\n", account1.Balance, account1.State)

	//find&update
	ee,_ := store.Find(account1.Id.GetGuid())
	restoredAccount1 := accounts.Load(account1.Id.Guid,ee)
	restoredAccount1.Handle(&accounts.WithdrawCommand{
		WithGuid:    es.WithGuid{Guid: es.NewGuid()},
		Amount:      50,
		Description: "gimme the money!",
	})
	// accountcreated, deposited,deposited
	//update withdrew
	fmt.Printf("Restored-Account1-> list uncommited changes:\n")
	restoredAccount1.ShowChanges()
	err= store.Update(restoredAccount1.Id.GetGuid(),restoredAccount1.Version-1,restoredAccount1.GetChanges())
	restoredAccount1.ClearChanges()
	if err != nil {
		fmt.Printf("Failed to update aggregate. \tError: %v",err)
	}else{
		fmt.Printf("Current Balance: %d\t State:%s",restoredAccount1.Balance,restoredAccount1.State)
	}
	restoredAccount1.ShowChanges()
}