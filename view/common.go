package view

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SendSucResp(c *gin.Context, res interface{}) {
	logrus.WithFields(logrus.Fields{
		"data": res,
	}).Info("suc resp")

	c.JSON(200, map[string]interface{}{
		"code": 0,
		"data": res,
		"msg":  "success",
	})
}

func SendErrResp(c *gin.Context, code int, res interface{}) {
	logrus.WithFields(logrus.Fields{
		"error": res,
	}).Warn("err resp")

	c.JSON(200, map[string]interface{}{
		"code": code,
		"data": res,
		"msg":  "error",
	})
}

func SendListResp(c *gin.Context, count int, res interface{}) {
	logrus.WithFields(logrus.Fields{
		"data": res,
		"count": count,
	}).Info("suc list resp")

	c.JSON(200, map[string]interface{}{
		"data":  res,
		"msg":   "success",
		"count": count,
	})
}
