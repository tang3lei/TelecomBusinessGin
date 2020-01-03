package model

type Employee struct {
	Base
	JobNumber string `json:"job_number" gorm:"column:job_number"`
	Name      string `json:"name" gorm:"column:name"`
	Password  string `json:"password" gorm:"column:password"`
}

func (Employee) TableName() string {
	return "employees"
}
