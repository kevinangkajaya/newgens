package test

import (
	"newgens/repository/mysql"
	"newgens/src"
	"testing"
)

func TestGetDataMt202(t *testing.T) {
	db := src.ConnectMysql()
	defer db.Close()

	mt202RepoMysql := mysql.NewRepoMt202Mysql(db)

	mt202, err := mt202RepoMysql.GetData()
	if err != nil {
		t.Fatalf("Get data failed: %s", err)
	}
	for _, v := range mt202 {
		t.Log(v)
	}
}
