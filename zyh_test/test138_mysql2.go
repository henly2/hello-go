package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root@localhost" //用户名
	dbpassword = "mysql"          //密码
	dbname     = "students"       //表名
)

func main() { //root:密码@tcp（127.0.0.1:3306）/数据库名？chartset=utf8
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/golang02?charset=utf8")
	fmt.Println(db)
	checkErr1(err)

	//插入数据
	stmt, err := db.Prepare("INSERT students SET name=?,age=?,high=?")
	checkErr1(err)

	res, err := stmt.Exec("码农", 1, 1.11)
	checkErr1(err)

	id, err := res.LastInsertId()
	checkErr1(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update students set name=? where id=?")
	checkErr1(err)

	res, err = stmt.Exec("码农二代", id)
	checkErr1(err)

	affect, err := res.RowsAffected()
	checkErr1(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("select * from students")
	checkErr1(err)

	for rows.Next() {
		var id int
		var name string
		var age int
		var high float64
		var gender string
		var cls_id interface{}
		//var high float64
		err = rows.Scan(&id, &name, &age, &high, &gender, &cls_id)
		//checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(age)
		fmt.Println(high)
		fmt.Println(gender)
		fmt.Println(cls_id)
	}

	//删除数据
	stmt, err = db.Prepare("delete from students where id=?")
	checkErr1(err)

	res, err = stmt.Exec(2)
	checkErr1(err)

	affect, err = res.RowsAffected()
	checkErr1(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr1(err error) {
	if err != nil {
		panic(err)
	}
}
