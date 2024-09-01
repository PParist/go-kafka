package repositories

import (
	"consumer/entities"
	"fmt"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepository {
	db.AutoMigrate(entities.Account{})
	return &accountRepository{db: db}
}

func (r *accountRepository) Save(account entities.Account) error {

	if exitRecode := r.db.Where("account_uuid = ?", account.AccountUUID).First(&account); exitRecode.Error != nil {
		if result := r.db.Save(&account); result.Error != nil {
			return result.Error
		}
		return nil
	} else {
		return gorm.ErrDuplicatedKey
	}

}

func (r *accountRepository) Update(account entities.Account) error {
	if result := r.db.Model(&entities.Account{}).Where("account_uuid = ?", account.AccountUUID).Update("balance", account.Balance); result.Error != nil {
		fmt.Printf("error : %v", result.Error)
		return result.Error
	}
	return nil
}

func (r *accountRepository) Delete(accountUUID string) error {
	if result := r.db.Where("account_uuid = ?", accountUUID).Delete(&entities.Account{}); result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *accountRepository) FindAll() (*[]entities.Account, error) {
	accounts := []entities.Account{}
	if result := r.db.Find(&accounts); result.Error != nil {
		return nil, result.Error
	}
	return &accounts, nil
}
func (r *accountRepository) FindByUUID(uuid string) (*entities.Account, error) {
	account := entities.Account{}
	if result := r.db.Where("account_uuid = ?", uuid).First(&account); result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}
