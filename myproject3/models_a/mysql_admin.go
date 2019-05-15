package models_a

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Admingrom struct {
	Root       string `json:"root"`
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Password   string `json:"password"`
	Registtime string `json:"registtime"`
}

var DB *gorm.DB

func init() {
	newdb, err := gorm.Open("mysql", "root:mysql@/mysql_orm?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("err=", err)
	}
	newdb.CreateTable(&Admingrom{}) //创建表
	DB = newdb
}
