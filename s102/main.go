package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	fmt.Println(" main function ...")

	// database sql
	db, err := sql.Open("mysql", "root:dg-mall.com@tcp(192.168.11.126:3306)/dg-mall-store?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}

