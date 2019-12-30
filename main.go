package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"upc.edu.cn/telecom_business/telecom_business_api/dal/db"
)
func main() {
	r := gin.Default()
	useUrlPattern(r)
	db.InitDb()
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}