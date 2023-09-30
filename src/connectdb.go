package src

import (
	"fmt"
	"newgens/config"

	"github.com/jmoiron/sqlx"
)

func ConnectMysql() *sqlx.DB {
	configs := config.GetConfig()
	dbCred := configs.Database

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbCred.UserName, dbCred.Password, dbCred.Host, dbCred.Port, dbCred.Name)
	// db, err := sql.Open("mysql", dataSourceName)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	return db
}
