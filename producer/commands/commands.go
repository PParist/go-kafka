package commands

type OpenAccountCommand struct {
	AccountHolder  string  `validate:"required"`
	AccountType    string  `validate:"required"`
	OpeningBalance float64 `validate:"required"`
}

type DepositFunCommand struct {
	AccountUUID string  `validate:"required,uuid"`
	Amount      float64 `validate:"required"`
}

type WithdrawFunCommand struct {
	AccountUUID string  `validate:"required,uuid"`
	Amount      float64 `validate:"required"`
}

type CloseAccountCommand struct {
	AccountUUID string `validate:"required,uuid"`
}
