package models

import (
	"fmt"
	u "github.com/cryptos-data-store/utils"

	"github.com/jinzhu/gorm"
)

type Crypto struct {
	gorm.Model
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
	UserId   uint    `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (crypto *Crypto) Validate() (map[string]interface{}, bool) {

	if crypto.Name == "" {
		return u.Message(false, "Crypto name should be on the payload"), false
	}

	if crypto.Symbol == "" {
		return u.Message(false, "Symbol should be on the payload"), false
	}

	if crypto.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (crypto *Crypto) Create() map[string]interface{} {

	if resp, ok := crypto.Validate(); !ok {
		return resp
	}

	GetDB().Create(crypto)

	resp := u.Message(true, "success")
	resp["crypto"] = crypto

	return resp
}

func Get(id uint) *Crypto {

	crypto := &Crypto{}

	err := GetDB().Table("cryptos").Where("id = ?", id).First(crypto).Error
	if err != nil {
		return nil
	}

	return crypto
}

func GetAll() []*Crypto {

	cryptos := make([]*Crypto, 0)

	err := GetDB().Find(&cryptos).Error
	if err != nil {
		return nil
	}

	return cryptos
}

func GetUserCrypto(user uint) []*Crypto {

	cryptos := make([]*Crypto, 0)

	err := GetDB().Table("cryptos").Where("user_id = ?", user).Find(&cryptos).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cryptos
}

func (crypto *Crypto) Update(newCrypto Crypto) map[string]interface{} {

	if resp, ok := crypto.Validate(); !ok {
		return resp
	}

	err := GetDB().Table("cryptos").Where("id = ?", newCrypto.ID).First(crypto).Error
	if err != nil {
		return nil
	}

	crypto = &newCrypto
	GetDB().Save(&crypto)

	resp := u.Message(true, "success")
	resp["crypto"] = crypto

	return resp
}

func Delete(id uint) *Crypto {

	crypto := &Crypto{}

	err := GetDB().Table("cryptos").Where("id = ?", id).First(crypto).Error
	if err != nil {
		return nil
	}

	db.Delete(crypto)

	return crypto
}
