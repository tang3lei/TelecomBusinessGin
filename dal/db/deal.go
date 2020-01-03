package db

import (
	"upc.edu.cn/telecom_business/telecom_business_api/model"
)

type DBDeal struct {
	Base
}

func Deal() *DBDeal {
	return &DBDeal{Base{tablename:model.Deal{}.TableName()}}
}
func (p *DBDeal) Transaction(tx *Transction) *DBDeal {
	p.transaction = tx
	return p
}

func (p *DBDeal) GetByMap(mp map[string]interface{}, out interface{}) error {
	rm := p.runtime()
	for k,v := range mp {
		rm = rm.Where(k,v)
	}
	rm = rm.Find(out)
	return rm.Error
}