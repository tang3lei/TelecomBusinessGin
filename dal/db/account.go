package db

import "upc.edu.cn/telecom_business/telecom_business_api/model"

type DBAccount struct {
	Base
}

func Account() *DBAccount {
	return &DBAccount{Base{tablename:model.Account{}.TableName()}}
}

func (p *DBAccount) Transaction(tx *Transction) *DBAccount {
	p.transaction = tx
	return p
}

func (p *DBAccount) GetByPhoneNumber(pn string,out interface{}) error {
	return p.Base.runtime().Table(p.tablename).Where("phone_number = ?",pn).First(out).Error
}
