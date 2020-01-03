package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func optionsMiddleWare(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Header("Access-Control-Allow-Headers", "e_name,xl_ck, xl_e_name,Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	c.Header("Access-Control-Expose-Headers", "e_name,xl_ck, xl_e_name, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Access-Control-Allow-Credentials")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}

func useMiddleWare(e *gin.Engine)  {
	e.Use(optionsMiddleWare)
}
