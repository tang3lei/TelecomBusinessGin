package model

import "time"

type Base struct {
	Id       int       `json:"id" gorm:"column:id"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

type Item interface {
	GetID() int
	Omit() []string
}

func (p Base) GetID() int {
	return p.Id
}

func (p Base) Omit() []string {
	return []string{"id","create_at","update_at"}
}