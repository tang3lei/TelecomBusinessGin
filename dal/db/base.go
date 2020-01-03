package db

import (
	"github.com/jinzhu/gorm"
	"upc.edu.cn/telecom_business/telecom_business_api/model"
)

type Base struct {
	tablename string
	transaction *Transction
}

func (p *Base) Transaction(tx *Transction) *Base {
	p.transaction = tx
	return p
}

func (p *Base) Debug() *gorm.DB {
	return p.runtime()
}

func (p *Base) runtime() *gorm.DB {
	if p.transaction != nil {
		return p.transaction.tx
	}
	return orm
}

func (p *Base) Get(base model.Item, out interface{}) error {
	return p.runtime().Where("id = ?", base.GetID()).First(out).Error
}

func (p *Base) ApplyQuery(arg *QueryArg, out interface{}) error {
	runtime := p.runtime()
	for i := range arg.argKey {
		runtime = runtime.Where(arg.argKey[i],arg.argValue[i])
	}
	runtime = runtime.Order(arg.Order)
	runtime = runtime.Limit(arg.Limit)
	runtime = runtime.Offset(arg.Offset)
	runtime = runtime.Find(out)

	return runtime.Error
}

func (p *Base) Create(val model.Item) error {
	runtime := p.runtime().Table(p.tablename)
	runtime = runtime.Create(val)
	return runtime.Error
}

func (p *Base) Update(val model.Item) error {
	runtime := p.runtime().Table(p.tablename).Model(val).Omit(val.Omit()...).Where("id = ?", val.GetID())
	runtime = runtime.Updates(val)
	return runtime.Error
}

func (p *Base) Delete(val model.Item) error {
	return p.runtime().Table(p.tablename).Model(val).Where("id = ?",val.GetID()).Delete(val).Error
}
