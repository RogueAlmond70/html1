package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

}
