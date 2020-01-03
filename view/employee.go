package view

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
	"upc.edu.cn/telecom_business/telecom_business_api/model"
)

type LoginReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

func LoginEmployee(c *gin.Context)  {
	var req LoginReq
	err := c.BindJSON(&req)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}
	var employee model.Employee
	err = db.Employee().FindByName(req.Name,&employee)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}
	if req.Password != employee.Password {
		SendErrResp(c,1,errors.New("error pwd"))
		return
	}
	c.Header("xl_ck","123456")
	c.Header("xl_e_name",employee.Name)
	SendSucResp(c, "suc")
}

func ListEmployee(c *gin.Context)  {
	arg,_ := NewQueryArg(c)
	arg.EqualInt("id","id")

	var res []model.Employee
	err := db.Employee().ApplyQuery(arg,&res)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}
	SendSucResp(c,res)
}

func UpdateEmployee(c *gin.Context)  {
	var employee model.Employee
	var err error
	err = c.BindJSON(&employee)
	if err != nil {
		SendErrResp(c,1,err)
		return
	}


	if employee.Id == 0 {
		employee.CreateAt = time.Now()
		employee.UpdateAt = employee.CreateAt
		err = db.Employee().Create(&employee)
	} else {
		err = db.Employee().Update(&employee)
	}
	if err != nil {
		SendErrResp(c,1,err)
		return
	}

	SendSucResp(c, employee)
}