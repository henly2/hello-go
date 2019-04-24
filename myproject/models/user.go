package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Userorm struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Password string `json:"password"`
}

func init() {
	//开启打印模式
	orm.Debug = true
	//连接数据库
	orm.RegisterDataBase("default", "mysql", "root:mysql@tcp(127.0.0.1:3306)/mysql_orm?charset=utf8")
	//映射model数据
	orm.RegisterModel(new(Userorm))
	//生成响应表
	orm.RunSyncdb("default", false, true)

	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	fmt.Println("Connect to redis error", err)
	//	return
	//}
	//defer c.Close()

}
