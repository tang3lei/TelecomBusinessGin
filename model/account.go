package model

type Account struct {
	Base
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number"`
	UserName    string  `json:"user_name" gorm:"column:user_name"`
	Status      int     `json:"status" gorm:"column:status"`
	Balance     float64 `json:"balance" gorm:"column:balance"`
	Package     string  `json:"package" gorm:"column:package"`
	Info        string  `json:"info" gorm:"column:info"`
	Desc2       string  `json:"desc2" gorm:"column:desc2"`
}

func (Account) TableName() string {
	return "accounts"
}
