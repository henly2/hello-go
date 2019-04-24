package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Student0 struct {
	id   int
	name string
	age  int
}

func main() {

	db, err := gorm.Open("mysql", "root:mysql@/golang04?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("err=", err)
	}
	//db.DropTable("students") //删除表

	type Li struct {
		ID        int    `gorm:"primary_key"`
		Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
		Ua        string `gorm:"type:varchar(256);not null;"`
		Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
		Hash      uint64 `gorm:"unique_index:hash_idx;"`
		CreatedAt time.Time
	}

	//if !db.HasTable(&Lik{}) {
	//	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Lik{}).Error; err != nil {
	//		panic(err)
	//	}
	//}
	//db.CreateTable(&Li{})
	//var li Li
	//db.Create(li) //?
	db.Model(&Li{}).AddIndex("idx_user_name", "Title")

}
