package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

const createUserTable = `CREATE TABLE users (
	id int(11) NOT NULL AUTO_INCREMENT,
		username varchar(100) NOT NULL,
		password text NOT NULL,
		PRIMARY KEY (id),
		UNIQUE KEY username (username)
	) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;`

func Seed() {
	dbName := viper.GetString("db_name")
	_, err := db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Println(fmt.Sprintf("database [%s] probably exists : error %s", dbName, err))
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		log.Println(fmt.Sprintf("error while using [%s] : error %s", dbName, err))
	}

	_, err = db.Exec(createUserTable)
	if err != nil {
		log.Println(fmt.Sprintf("table [users] probably exists : error %s", err))
	}

}
