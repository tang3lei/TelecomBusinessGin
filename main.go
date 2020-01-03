package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
)
func main() {
	r := gin.Default()
	useMiddleWare(r)
	useUrlPattern(r)
	db.InitDb()
	err := r.Run()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Info("init server")
	}
}