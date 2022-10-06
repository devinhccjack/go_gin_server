package Mysql

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
)

var DB *gorm.DB

func init()  {
    var err error
	dsn := "root:root123456@tcp(127.0.0.1:3306)/testdb"

	// "wuyu:MIDSUMMERfish0@/gin?charset=utf8&parseTime=True&loc=Local"
    DB, err = gorm.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }
    if DB.Error != nil {
        fmt.Printf("database error %v", DB.Error)
    }
}