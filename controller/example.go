package controller

import (
	"context"
	"log"
	dbcon "rest-example-env/dbconn"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	Sex  string
	age  int
}

func (co *Controller) GetExample(c *gin.Context) {
	conn := co.NewConnection()
	db := dbcon.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)

	defer db.Close()
	ctx := context.Background()

	// Execute query
	rows, err := db.QueryContext(ctx, "SELECT name,sex,age FROM person")
	if err != nil {
		log.Println("Error Query: ", err.Error())
		return
	}

	defer rows.Close()
}
