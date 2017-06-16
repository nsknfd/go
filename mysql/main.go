package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	CheckErr(err)
	defer db.Close()
	// 插入数据
	stmt, err := db.Prepare("INSERT INTO user(id, name) values(?, ?)")
	CheckErr(err)

	res, err := stmt.Exec(os.Args[1], os.Args[2])
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	fmt.Println(id)
}
