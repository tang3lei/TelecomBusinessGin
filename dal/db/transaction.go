package db

import "github.com/jinzhu/gorm"

type Transction struct {
	tx *gorm.DB
}

func BeginTransction() *Transction {
	tx := orm.Begin()
	return &Transction{tx}
}

func (p *Transction) Commit() {
	p.tx.Commit()
}

func (p *Transction) Rollback() {
	p.tx.Rollback()
}
