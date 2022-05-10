package dbconn

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectDB(server string, port string, user string, password string, database string) *sqlx.DB {
	var db *sqlx.DB
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)

	db, err = sqlx.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	return db

}
