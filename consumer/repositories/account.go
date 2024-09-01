package repositories

import "consumer/entities"

type AccountRepository interface {
	Save(account entities.Account) error
	Update(account entities.Account) error
	Delete(accountUUID string) error
	FindAll() (*[]entities.Account, error)
	FindByUUID(uuid string) (*entities.Account, error)
}
