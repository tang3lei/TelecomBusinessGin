package db

import "upc.edu.cn/telecom_business/telecom_business_api/model"

type DBAccount struct {
	Base
}

func Account() *DBAccount {
	return &DBAccount{Base{tablename:model.Account{}.TableName()}}
}
