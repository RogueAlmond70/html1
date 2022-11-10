package Database

import (
	"Spotlas/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func getConn() string {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		viper.GetString(config.DBHost),
		viper.GetInt64(config.DBPort),
		viper.GetString(config.DBUser),
		viper.GetString(config.DBPass),
		viper.GetString(config.DBName))
	return psqlconn
}

func DBSetup() *sql.DB {
	connStr := getConn()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Check(db *sql.DB) {
	rows, err := db.Query(config.Check1)
	defer rows.Close()
	for rows.Next() {
		var data sql.NullString

		err = rows.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)
	}
}

func Question1(db *sql.DB) {
	rows, err := db.Query(config.Question1)
	defer rows.Close()
	for rows.Next() {
		var coordinates sql.NullString

		err = rows.Scan(&coordinates)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(coordinates)
	}

	if err != nil {
		log.Fatal(err)
	}

}

func Q1(db *sql.DB) {
	rows, err := db.Query(config.StripPrefix)
	defer rows.Close()
	for rows.Next() {
		var data sql.NullString

		err = rows.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)
	}

	if err != nil {
		log.Fatal(err)
	}

}
