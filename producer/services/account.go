package services

import "producer/commands"

type AccountService interface {
	OpenAccount(command commands.OpenAccountCommand) error
	Deposit(command commands.DepositFunCommand) error
	Withdraw(command commands.WithdrawFunCommand) error
	CloseAccount(command commands.CloseAccountCommand) error
}
