package view

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
	"upc.edu.cn/telecom_business/telecom_business_api/model"
	"upc.edu.cn/telecom_business/telecom_business_api/tools"
)

func ListAccount(c *gin.Context) {
	arg, _ := NewQueryArg(c)
	arg.Equal("phone_number", "phone_number").
		EqualInt("id", "id")

	var res []model.Account
	err := db.Account().ApplyQuery(arg, &res)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	SendSucResp(c, res)
}

func UpdateAccount(c *gin.Context) {
	var account model.Account
	var err error
	err = c.BindJSON(&account)
	if err != nil {
		SendErrResp(c, 1, err)
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
		SendErrResp(c, 1, err)
		return
	}

	SendSucResp(c, account)
}

func DeleteAccount(c *gin.Context) {
	id := tools.GetQueryInt64(c, "id", 0)
	if id == 0 {
		SendErrResp(c, 1, errors.New("request need id"))
		return
	}
	var res model.Account
	res.Id = int(id)
	err := db.Account().Delete(&res)
	if err != nil {
		SendErrResp(c, 1, err)
		return
	}
	SendSucResp(c, "success delete")
}

func UpdateAccountBalance(c *gin.Context) {
	phoneNumber := tools.GetQueryString(c, "phone_number", "")
	addNumber := tools.GetQueryFloat64(c, "add_number", 0)
	if phoneNumber == "" || phoneNumber == "undefined" {
		SendErrResp(c, 1, errors.New("request need phone_number"))
		return
	}
	action := db.BeginTransction()
	var account model.Account
	err := db.Account().Transaction(action).GetByPhoneNumber(phoneNumber, &account)
	if err != nil {
		SendErrResp(c, 1, err)
		action.Rollback()
		return
	}
	account.Balance += addNumber
	err = db.Account().Transaction(action).Update(&account)
	if err != nil {
		SendErrResp(c, 1, err)
		action.Rollback()
		return
	}
	now := time.Now()
	deal := model.Deal{
		DealName:"余额充值",
		DealTime:now.Unix(),
		PhoneNumber:phoneNumber,
		JobName:c.GetHeader("e_name"),
		Type:-1,
		Cost:addNumber,
	}
	deal.CreateAt = now
	deal.UpdateAt = now
	err = db.Deal().Transaction(action).Create(&deal)
	if err != nil {
		SendErrResp(c, 1, err)
		action.Rollback()
		return
	}
	action.Commit()
	SendSucResp(c, account)
}

func UseAccountBalance(c *gin.Context) {
	phoneNumber := tools.GetQueryString(c, "phone_number", "")
	if phoneNumber == "" || phoneNumber == "undefined" {
		SendErrResp(c, 1, errors.New("request need phone_number"))
		return
	}
	kuandai := tools.GetQueryString(c, "kuandai", "")
	liuliang := tools.GetQueryString(c, "liuliang", "")
	k := strings.Split(kuandai, "-")
	l := strings.Split(liuliang, "-")

	var account model.Account
	action := db.BeginTransction()
	err := db.Account().Transaction(action).GetByPhoneNumber(phoneNumber, &account)
	if err != nil {
		SendErrResp(c, 1, err)
		action.Rollback()
		return
	}
	now := time.Now()
	deal := model.Deal{
		DealTime:now.Unix(),
		PhoneNumber:phoneNumber,
		JobName:c.GetHeader("e_name"),
	}
	deal.CreateAt = now
	deal.UpdateAt = now
	var resNum float64
	if len(k) > 1 {
		resNum = float64(tools.StringGetInt(k[2]))
		deal.DealName = "宽带缴费"
		deal.Type = 1
		deal.Cost = -resNum
	}else if len(l) > 1 {
		resNum = float64(tools.StringGetInt(l[2]))
		deal.DealName = "流量缴费"
		deal.Type = 2
		deal.Cost = -resNum
	}else {
		SendErrResp(c, 1, errors.New("empty param"))
		action.Rollback()
		return
	}
	account.Balance -= resNum
	if account.Balance < 0 {
		SendErrResp(c,1,"余额不足")
		action.Rollback()
		return
	}
	err = db.Account().Transaction(action).Update(&account)
	if err != nil {
		SendErrResp(c, 1, err)
		action.Rollback()
		return
	}

	err = db.Deal().Transaction(action).Create(&deal)
	if err != nil {
		SendErrResp(c, 1, err)
		action.Rollback()
		return
	}
	action.Commit()
	SendSucResp(c, "suc")
}
