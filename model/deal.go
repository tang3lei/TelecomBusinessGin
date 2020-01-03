package model

type Deal struct {
	Base
	DealName    string  `json:"deal_name" gorm:"column:deal_name"`
	DealTime    int64   `json:"deal_time" gorm:"column:deal_time"`
	PhoneNumber string  `json:"phone_number" gorm:"column:phone_number"`
	JobName   string  `json:"job_name" gorm:"column:job_name"`
	Type        int     `json:"type" gorm:"column:type"` // -1 充值 1 宽带 2 流量 3 套餐
	Cost        float64 `json:"cost" gorm:"column:cost"`
}

func (Deal) TableName() string {
	return "deals"
}
