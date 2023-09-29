package main

import (
	"fmt"
	"newgens/models"
	"newgens/repository/mysql"
	"newgens/src"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertDataMt202(t *testing.T) {
	db := src.ConnectMysql()
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	var data models.MT202
	if err := mt202RepoMysql.InsertData(&data); err != nil {
		t.Errorf("Insert data failed: %s", err)
	}
}

func TestGetDataMt202(t *testing.T) {
	db := src.ConnectMysql()
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	mt202, err := mt202RepoMysql.GetData()
	if err != nil {
		t.Errorf("Get data failed: %s", err)
	}
	for _, v := range mt202 {
		fmt.Println(v)
	}
}
