package main

import (
	"database/sql"
	"todolist/functions"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//sql.Open("<banco de dados>","<user>:<password>@tcp(localhost:<porta>)/<nome do DB")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/todolist")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	if err != nil {
		panic(err)
	}

	functions.AppToDoList(db)

}
