package main

import (
	"github.com/gin-gonic/gin"
	"upc.edu.cn/telecom_business/telecom_business_api/view"
)

var URL_PATTERN = map[string]gin.HandlersChain{
	"api/account/list": {view.ListAccount},
	"api/account/update": {view.UpdateAccount},
}

func useUrlPattern(p *gin.Engine)  {
	for url, handlers := range URL_PATTERN{
		p.GET(url,handlers...)
		p.POST(url,handlers...)
	}
}
