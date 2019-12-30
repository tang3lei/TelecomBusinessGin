package view

import "github.com/gin-gonic/gin"

func SendSucResp(c *gin.Context, res interface{}) {
	c.JSON(200, map[string]interface{}{
		"code": 0,
		"data": res,
		"msg":  "success",
	})
}

func SendErrResp(c *gin.Context, code int, res interface{}) {
	c.JSON(200, map[string]interface{}{
		"code": code,
		"data": res,
		"msg":  "error",
	})
}

func SendListResp(c *gin.Context, count int, res interface{}) {
	c.JSON(200, map[string]interface{}{
		"data":  res,
		"msg":   "success",
		"count": count,
	})
}
