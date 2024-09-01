package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(OpenAccountEvent{}).Name(),
	reflect.TypeOf(DepositFunEvent{}).Name(),
	reflect.TypeOf(WithdrawFunEvent{}).Name(),
	reflect.TypeOf(CloseAccountEvent{}).Name(),
}

type Event interface {
}

type OpenAccountEvent struct {
	AccountUUID    string  `json:"account_uuid"`
	AccountHolder  string  `json:"account_holder"`
	AccountType    string  `json:"account_type"`
	OpeningBalance float64 `json:"opening_balance"`
}

type DepositFunEvent struct {
	AccountUUID string  `json:"account_uuid"`
	Amount      float64 `json:"amount"`
}

type WithdrawFunEvent struct {
	AccountUUID string  `json:"account_uuid"`
	Amount      float64 `json:"amount"`
}

type CloseAccountEvent struct {
	AccountUUID string `json:"account_uuid"`
}
