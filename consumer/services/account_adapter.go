package services

import (
	"consumer/entities"
	"consumer/repositories"
	"encoding/json"
	"events"
	"fmt"
	"log"
	"reflect"
)

type accountEventHandler struct {
	repo repositories.AccountRepository
}

func NewAccountEventHandler(repo repositories.AccountRepository) EventHandler {
	return &accountEventHandler{repo: repo}
}

func (s *accountEventHandler) Handle(topic string, eventBytes []byte) {

	if !json.Valid(eventBytes) {
		log.Println("Invalid JSON received")
		return
	}

	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("value event %#v\n", event)
		bankaccount := entities.Account{
			AccountUUID:   event.AccountUUID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}

		if err = s.repo.Save(bankaccount); err != nil {
			log.Println(err)
			return
		}
		log.Printf("event : %#v", event)
	case reflect.TypeOf(events.DepositFunEvent{}).Name():
		event := events.DepositFunEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := s.repo.FindByUUID(event.AccountUUID)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance += event.Amount
		if err = s.repo.Update(*bankAccount); err != nil {
			log.Println(err)
			return
		}
		log.Printf("event : %#v", event)
	case reflect.TypeOf(events.WithdrawFunEvent{}).Name():
		event := events.WithdrawFunEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := s.repo.FindByUUID(event.AccountUUID)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance -= event.Amount
		if err = s.repo.Update(*bankAccount); err != nil {
			log.Println(err)
			return
		}
		log.Printf("event : %#v", event)
	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Println(err)
			return
		}
		if err = s.repo.Delete(event.AccountUUID); err != nil {
			log.Println(err)
			return
		}
		log.Printf("event : %#v", event)
	default:
		log.Println("no event handler")
		return
	}
}
