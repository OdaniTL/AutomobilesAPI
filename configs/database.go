package configs

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var ones sync.Once

func NewConnection() (*sql.DB, error) {
	var DB *sql.DB
	var errr error
	ones.Do(func() {
		//ssmode=disable
		sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("P_SQL_HOST"),
			os.Getenv("P_SQL_PORT"),
			os.Getenv("P_SQL_USER"),
			os.Getenv("P_SQL_PASSWORD"),
			os.Getenv("P_SQL_DBNAME"))
		db, err1 := sql.Open("postgres", sqlInfo)

		err2 := db.Ping()
		errr = errors.Join(err1, err2)
		DB = db
	})
	return DB, errr
}
