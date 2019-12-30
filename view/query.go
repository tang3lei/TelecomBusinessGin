package view

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
)

func NewQueryArg(c *gin.Context) (*db.QueryArg, error) {
	arg := db.QueryArg{}
	err := c.Request.ParseMultipartForm(8 << 20)
	arg.ParamsMap = c.Request.Form
	if limit, ok := arg.ParamsMap["limit"]; ok {
		l, _ := strconv.Atoi(limit[0])
		arg.Limit = l
	} else {
		arg.Limit = 20
	}
	if offset, ok := arg.ParamsMap["offset"]; ok {
		o, _ := strconv.Atoi(offset[0])
		arg.Offset = o
	} else {
		arg.Offset = 0
	}
	if order, ok := arg.ParamsMap["order"]; ok {
		arg.Order = order[0]
	} else {
		arg.Order = "id"
	}
	return &arg, err
}
