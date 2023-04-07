package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 建立数据库连接
	db, err := sql.Open("mysql", "root:Wa@123456@tcp(172.20.8.110:31002)/fedx")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 执行查询
	rows, err := db.Query("SELECT id, name FROM project")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 遍历结果集
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id=%d, name=%s\n", id, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 执行插入
	//stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()
	//
	//res, err := stmt.Exec("Alice")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//lastID, err := res.LastInsertId()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("last inserted id=%d\n", lastID)
}
