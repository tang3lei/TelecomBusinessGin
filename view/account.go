package view

import (
	"github.com/gin-gonic/gin"
	"time"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
	"upc.edu.cn/telecom_business/telecom_business_api/model"
)



func ListAccount(c *gin.Context)  {
	arg,_ := NewQueryArg(c)
	arg.Equal("phone_number","phone_number").
		EqualInt("id","id")

	var res []model.Account
	err := db.Account().ApplyQuery(arg,&res)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}
	SendSucResp(c,res)
}

func UpdateAccount(c *gin.Context)  {
	var account model.Account
	var err error
	err = c.BindJSON(&account)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}


	if account.Id == 0 {
		account.CreateAt = time.Now()
		account.UpdateAt = account.CreateAt
		err = db.Account().Create(&account)
	} else {
		err = db.Account().Update(&account)
	}
	if err != nil {
		SendErrResp(c,1,err)
		return
	}

	SendSucResp(c,account)
}