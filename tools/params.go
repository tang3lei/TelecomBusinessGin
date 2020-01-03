package tools

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQueryString(c *gin.Context, key, defaultV string) string {
	return c.DefaultQuery(key, defaultV)
}

func GetQueryInt64(c *gin.Context, key string, defaultV int64) int64 {
	str := c.DefaultQuery(key, "")
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil || str == "" {
		return defaultV
	}
	return val
}

func GetQueryFloat64(c *gin.Context, key string, defaultV float64) float64 {
	str := c.DefaultQuery(key, "")
	val, err := strconv.ParseFloat(str, 64)
	if err != nil || str == "" {
		return defaultV
	}
	return val
}
