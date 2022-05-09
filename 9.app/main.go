package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	c := mysql.Config{
		User:      "root",
		Passwd:    "1234",
		Addr:      "localhost:3306",
		Net:       "tcp",
		DBName:    "phone_book", // contacts
		ParseTime: true,
	}

	fmt.Println(c.FormatDSN()) // db info

	db, err := sql.Open("mysql", c.FormatDSN()) //umesto c.formatdns moze se odmah napisati info za bazu ("mysql, user:pass@tcp(127.0.0.1:3306)/db_name")
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	phone_book := make(map[string]int) //made empty array of contact object - phone book

	var running = true
	var input int = 0
	var name string
	var phone int

	const (
		ADD_CONTACT    = 1
		REMOVE_CONTACT = 2
		VIEW_CONTACT   = 3
		EXIT           = 0
	)
	fmt.Print("Welcome to phone adress book app\n")
	for running {
		fmt.Println("[1]- Add contact")
		fmt.Println("[2]- Remove contact")
		fmt.Println("[3]- View contacts")
		fmt.Println("[0]- Exit\n---(Navigate by typing index)---")

		fmt.Scanln(&input)
		switch {
		case input == EXIT:
			running = false // break statement
		case input == ADD_CONTACT:
			fmt.Println("Enter name and phone number: ")
			fmt.Scanln(&name, &phone)
			phone_book[name] = phone // for key name assign value of variable phone
			row := db.QueryRowContext(context.Background(), "INSERT INTO contacts (name, phone) VALUES(?, ?)", name, phone)
			if err := row.Err(); err != nil {
				fmt.Println("db.QueryRowContext", err)
				return
			}
		case input == REMOVE_CONTACT:
			var del int
			i := 1 // index number
			for key, value := range phone_book {
				fmt.Printf("%v. %v: %v\n", i, key, value)
				i++
			}
			fmt.Println("Type number of the contact you want to delete: ")
			fmt.Scanln(&del)
			for key, value := range phone_book {
				if del == value { // value
					delete(phone_book, key) // first argument is the name of the map | the second one is key
				}
			}
			row := db.QueryRowContext(context.Background(), "DELETE FROM contacts WHERE phone=?", del)
			if err := row.Err(); err != nil {
				fmt.Println("db.QueryRowContext", err)
				return
			}
		case input == VIEW_CONTACT:
			i := 1
			results, err := db.Query("SELECT * FROM contacts")
			if err != nil {
				log.Fatal("Error message: ", err)
			}
			for results.Next() {
				err = results.Scan(&name, &phone)
				phone_book[name] = phone
				if err != nil {
					log.Fatal("error", err)
				}
				fmt.Printf("%v. %v: %v\n", i, name, phone)
				i++
			}
			if len(phone_book) == 0 {
				fmt.Println("Contact list is empty. Add a contact with option 1.")
				//running = false
				continue // returns to the welcome page with notice to add contacts :D
			}

		default:
			fmt.Println("Invalid input")
			continue
		}
	}
}
