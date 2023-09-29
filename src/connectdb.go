package src

import (
	"database/sql"
	"fmt"
	"newgens/config"
)

func ConnectMysql() *sql.DB {
	configs := config.GetConfig()
	dbCred := configs.Database

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbCred.UserName, dbCred.Password, dbCred.Host, dbCred.Port, dbCred.Name)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
