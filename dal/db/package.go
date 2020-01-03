package db

import "upc.edu.cn/telecom_business/telecom_business_api/model"

type DBPackage struct {
	Base
}

func Package() *DBPackage {
	return &DBPackage{Base{tablename:model.Package{}.TableName()}}
}

