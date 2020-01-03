package db

import "upc.edu.cn/telecom_business/telecom_business_api/model"

type DBEmployee struct {
	Base
}

func Employee() *DBEmployee {
	return &DBEmployee{Base{tablename:model.Employee{}.TableName()}}
}

func (p *DBEmployee) FindByName(jn string, out interface{}) error {
	return p.runtime().Table(p.tablename).Where("name = ?",jn).First(out).Error
}
