package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

const (
	createUserTable = "CREATE TABLE `users` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`username` varchar(100) NOT NULL," +
		"`password` text NOT NULL," +
		"PRIMARY KEY (`id`)," +
		"UNIQUE KEY `username` (`username`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8;"

	createMessageTable = "CREATE TABLE `messages` (" +
		"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
		"`username` varchar(100) DEFAULT NULL," +
		"`message` text," +
		"`timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;"
)

func seed() {
	fmt.Println("seeding db...")

	dbName := viper.GetString("db_name")
	_, err := db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Println(fmt.Sprintf("database [%s] probably exists: %s", dbName, err))
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		log.Println(fmt.Sprintf("error while using database [%s] : error %s", dbName, err))
	}

	_, err = db.Exec(createUserTable)
	if err != nil {
		log.Println(fmt.Sprintf("table [users] probably exists: %s", err))
	}

	_, err = db.Exec(createMessageTable)
	if err != nil {
		log.Println(fmt.Sprintf("table [messages] probably exists: %s", err))
	}

}
