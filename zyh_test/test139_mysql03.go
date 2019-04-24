package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/golang04?charset=utf8")
	if err != nil {
		fmt.Println("err=", err)
	}
	//插入数据
	stmt, err := db.Prepare("INSERT into students(name,age)values(?,?)")

	res, err := stmt.Exec("zs", 1.11)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//查询数据
	rows, err := db.Query("select * from golang04.students")
	for rows.Next() {
		var id int
		var name string
		var age int

		err = rows.Scan(&id, &name, &age)
		checkErr(err)

		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(age)
	}
	//更新数据
	stmt, err = db.Prepare("update golang04.students set name=? where id=?")
	res, err = stmt.Exec("码农", 2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	//删除数据
	stmt, err = db.Prepare("delete from golang04.students where id=?")
	checkErr(err)
	res, err = stmt.Exec(3)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
