package main

import (
	"fmt"
	"newgens/src"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db := ConnectMysql()
	// defer db.Close()

	// mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	// main use
	path := src.ReadConsole()
	data, err := src.ReadLines(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
