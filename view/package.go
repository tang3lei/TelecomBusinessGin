package view

import (
	"github.com/gin-gonic/gin"
	"time"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
	"upc.edu.cn/telecom_business/telecom_business_api/model"
)

func ListPackage(c *gin.Context)  {
	arg,_ := NewQueryArg(c)
	arg.EqualInt("id","id")

	var res []model.Package
	err := db.Package().ApplyQuery(arg,&res)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}
	SendSucResp(c,res)
}

func UpdatePackage(c *gin.Context)  {
	var pkg model.Package
	var err error
	err = c.BindJSON(&pkg)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}


	if pkg.Id == 0 {
		pkg.CreateAt = time.Now()
		pkg.UpdateAt = pkg.CreateAt
		err = db.Package().Create(&pkg)
	} else {
		err = db.Package().Update(&pkg)
	}
	if err != nil {
		SendErrResp(c,1,err)
		return
	}

	SendSucResp(c, pkg)
}