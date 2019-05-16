package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserSqlite struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("sqlite3", "testsqlite.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&UserSqlite{})

	// 创建
	user := UserSqlite{Name: "zhangsan", Age: 18}
	db.NewRecord(user) // => 主键为空返回`true`

	db.Create(&user)

	//查询1
	err = db.Where("Name = ?", "zhangsan").Find(&user).Error
	if err != nil {
		fmt.Println("ss")
		return
	}
	//查询2
	err = db.Model(&UserSqlite{Name: "zhangsan", Age: 18}).Find(&user).Error
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	//查询3
	db.Where("Name LIKE?", "%zh%").Find(&user)
	//查询4
	err = db.Model(&user).Where("Age = ?", 18).Error
	if err != nil {
		return
	}

	//更新数据
	db.First(&user)
	user.Name = "ls"
	user.Age = 100
	db.Save(&user)

}
