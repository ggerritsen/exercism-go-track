// Package account provides functionality around bank accounts
package account

import "sync"

// Account is a bank account
type Account interface {
	// Close closes the bank acount
	Close() (payout int64, ok bool)
	// Balance returns the balance of the account
	Balance() (balance int64, ok bool)
	// Deposit adds (or withdraws in case of negative amount) money to the account
	Deposit(amount int64) (newBalance int64, ok bool)
}

type myAccount struct {
	sync.RWMutex
	balance int64
	closed bool
}

func (a *myAccount) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	if a.balance < 0 {
		return 0, false
	}

	bal := a.balance
	a.balance = 0
	a.closed = true
	return bal, true
}

func (a *myAccount) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()
	return a.balance, !a.closed
}

func (a *myAccount) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return a.balance, false
	}

	newBal := a.balance + amount
	if newBal < 0 {
		// too high withdrawal
		return a.balance, false
	}

	a.balance = newBal
	return a.balance, true
}

// Open opens an account with an initial deposit
func Open(initialDeposit int64) Account {
	if initialDeposit < 0 {
		return nil
	}
	return &myAccount{balance: initialDeposit}
}

