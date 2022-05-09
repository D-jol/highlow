package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Drivers: ", sql.Drivers())
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/phonebook")

	if err != nil {
		log.Fatal("Non succesful connection to db.")
	}
	time.Sleep(1 * time.Second)
	defer db.Close()
}
