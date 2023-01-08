package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"database/sql"
	"fmt"
	"log"
)


var db *sql.DB


func main() {
	var err error

	// setup viper
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(viper.GetString("DBUSER"))

	// capture connection properties
	cfg := mysql.Config{
		User: viper.GetString("DBUSER"),
		Passwd: viper.GetString("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: viper.GetString("DBNAME"),
		AllowNativePasswords: true,
	}

	// Get a db handle
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}
