package main

import (
	"fmt"
	"newgens/config"
	"newgens/models"
	"newgens/repository/mysql"
	"newgens/src"
	"strings"
)

func main() {
	configs := config.GetConfig()

	db := src.ConnectMysql(configs)
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	// main use
	path := src.ReadConsole()
	data, err := src.ReadLines(path)
	if err != nil {
		panic(err)
	}

	dataReady, err := models.NewMt202RawFromFile(strings.Join(data, "\n"))
	if err != nil {
		panic(err)
	}

	mt202, err := models.NewMT202FromRaw(dataReady)
	if err != nil {
		panic(err)
	}
	fmt.Println("convert to table success")

	if mt202 != nil {
		err = mt202RepoMysql.InsertData(mt202)
		if err != nil {
			panic(err)
		}
		fmt.Println("insert success")
	}
}
