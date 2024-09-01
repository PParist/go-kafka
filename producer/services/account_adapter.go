package services

import (
	"events"
	"producer/commands"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return &accountService{eventProducer: eventProducer}
}

func (s *accountService) OpenAccount(command commands.OpenAccountCommand) error {

	if err := validate.Struct(&command); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}

	event := events.OpenAccountEvent{
		AccountUUID:    uuid.New().String(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}

	if err := s.eventProducer.Produce(event); err != nil {
		return err
	}

	return nil
}

func (s *accountService) Deposit(command commands.DepositFunCommand) error {

	if err := validate.Struct(&command); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	event := events.DepositFunEvent{
		AccountUUID: command.AccountUUID,
		Amount:      command.Amount,
	}

	if err := s.eventProducer.Produce(event); err != nil {
		return err
	}

	return nil
}

func (s *accountService) Withdraw(command commands.WithdrawFunCommand) error {
	if err := validate.Struct(&command); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	event := events.WithdrawFunEvent{
		AccountUUID: command.AccountUUID,
		Amount:      command.Amount,
	}

	if err := s.eventProducer.Produce(event); err != nil {
		return err
	}

	return nil
}

func (s *accountService) CloseAccount(command commands.CloseAccountCommand) error {
	if err := validate.Struct(&command); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	event := events.CloseAccountEvent{
		AccountUUID: command.AccountUUID,
	}
	if err := s.eventProducer.Produce(event); err != nil {
		return err
	}
	return nil
}
