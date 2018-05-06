package dao

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
)

var (
	db *sql.DB
)

func init() {
	fmt.Println("init Db...")

	var err error
	db, err = sql.Open(viper.GetString("db_driver"),
		fmt.Sprintf("%s:%s@tcp(%s:3306)/",
			viper.GetString("db_user"),
			viper.GetString("db_pass"),
			viper.GetString("db_ip")))
	if err != nil {
		panic(err)
	}

	seed()
}
