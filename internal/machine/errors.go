package machine

import "errors"

var (
	ErrTransactionNotValid = errors.New("transaction not valid")
	ErrMachineBusyNow      = errors.New("machine busy now try latter")
	ErrProductNotFound     = errors.New("product not found")
	ErrProductRunningOut   = errors.New("product running out")
	ErrNotEnoughMoney      = errors.New("not enough money")
	ErrNotExistForDispense = errors.New("not exist for dispense")
)
