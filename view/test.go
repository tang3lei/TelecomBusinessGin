package view

import "github.com/gin-gonic/gin"

func TestFunc(c *gin.Context)  {
	c.JSON(1,"success")
}
