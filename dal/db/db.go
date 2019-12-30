package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	//_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"net/url"
)

var (
	orm *gorm.DB
	err error
)

func InitDb() {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("xl_dev", "123456asdA"),
		Host:   fmt.Sprintf("%s:%d", "121.199.78.126", 1433),
	}
	orm, err = gorm.Open("mssql", u.String())
	if err != nil {
		panic(err)
	}
	orm.LogMode(true)
}
