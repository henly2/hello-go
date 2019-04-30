package models_a

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Userlog struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Userlogintime string `json:"userlogintime"`
}

var Db *gorm.DB

func init() {
	newdb, err := gorm.Open("mysql", "root:mysql@/mysql_orm?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("err=", err)
	}
	newdb.CreateTable(&Userlog{}) //创建表
	Db = newdb
}
