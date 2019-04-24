package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增

	CreditCard CreditCard // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails     []Email    // One-To-Many (拥有多个 - Email表的UserID作外键)

	BillingAddress   Address // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // 忽略这个字段
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

// 重设列名
type Animal struct {
	AnimalId int64     `gorm:"column:beast_id"`         // 设置列名为`beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"` // 设置列名为`day_of_the_beast`
	Age      int64     `gorm:"column:age_of_the_beast"` // 设置列名为`age_of_the_beast`
}
type Teacher0 struct {
	gorm.Model
	Name string
	Age  int
}

//设置主键
type Animal0 struct {
	AnimalId int64 `gorm:"primary_key"` // 设置AnimalId为主键
	Name     string
	Age      int64
}

func main() {

	db, err := gorm.Open("mysql", "root:mysql@/golang05?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("err=", err)
	}
	db.CreateTable(&User{})     //创建表
	db.CreateTable(&Teacher0{}) //创建表

	db.Model(&User{}).ModifyColumn("age", "test") //修改age的数据类型为test

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	db.NewRecord(user) // => 主键为空返回`true`

	db.Create(&user)

	db.NewRecord(user)

}
