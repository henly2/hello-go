package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func testSelect() {

	db, err := sql.Open("mysql", "root:root@localhost(127.0.0.1:3306)/Test?charset=utf8")

	if err != nil {

		fmt.Printf("connect err")

	}

	rows, err1 := db.Query("select userid,username from tb_user limit 0,5")

	if err1 != nil {

		fmt.Println(err1.Error())

		return

	}

	defer rows.Close()

	fmt.Println("")

	cols, _ := rows.Columns()

	for i := range cols {

		fmt.Print(cols[i])

		fmt.Print("\t")

	}

	fmt.Println("")

	var userid int

	var username string

	for rows.Next() {

		if err := rows.Scan(&userid, &username); err == nil {

			fmt.Print(userid)
			fmt.Print("\t")

			fmt.Print(username)

			fmt.Print("\t\r\n")

		}

	}

}

func main() {

	testSelect()

}
