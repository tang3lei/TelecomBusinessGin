package main

import (
	"github.com/gin-gonic/gin"
	"upc.edu.cn/telecom_business/telecom_business_api/view"
)

var URL_PATTERN = map[string]gin.HandlersChain{
	"api/account/list":           {view.ListAccount},
	"api/account/update":         {view.UpdateAccount},
	"api/account/delete":         {view.DeleteAccount},
	"api/account/change_balance": {view.UpdateAccountBalance},
	"api/account/use_balance":    {view.UseAccountBalance},

	"api/deal/list":   {view.ListDeal},
	"api/deal/update": {view.UpdateDeal},
	"api/deal/month":  {view.MonthlyDeals},
	"api/deal/daily":  {view.DailyDeals},
	"api/deal/user":   {view.UserDailyDeals},


	"api/employee/list":   {view.ListEmployee},
	"api/employee/update": {view.UpdateEmployee},
	"api/employee/login":  {view.LoginEmployee},

	"api/package/list":   {view.ListPackage},
	"api/package/update": {view.UpdatePackage},
}

func useUrlPattern(p *gin.Engine) {
	for url, handlers := range URL_PATTERN {
		p.GET(url, handlers...)
		p.POST(url, handlers...)
	}
}
