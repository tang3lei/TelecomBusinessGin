package model

type Package struct {
	Base
	Name        string  `json:"name" gorm:"column:name"`
	Type        int     `json:"type" gorm:"column:type"`
	MonthlyCost float64 `json:"monthly_cost" gorm:"column:monthly_cost"`
	DailyCost   float64 `json:"daily_cost" gorm:"column:daily_cost"`
	Desc2       string  `json:"desc2" gorm:"column:desc2"`
}

func (Package) TableName() string {
	return "packages"
}
