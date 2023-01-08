package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?;", name)
	if err != nil {
		return nil, fmt.Errorf("abumsByArtist %q: %v", name, err)
	}

	defer rows.Close()
	// put rows in the output array
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

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

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}
