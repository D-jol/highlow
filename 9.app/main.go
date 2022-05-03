package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {

	c := mysql.Config{
		User:      "root",
		Passwd:    "1234",
		Addr:      "localhost:3306",
		Net:       "tcp",
		DBName:    "phone_book",
		ParseTime: true,
	}

	fmt.Println(c.FormatDSN()) // some info

	db, err := sql.Open("mysql", c.FormatDSN())
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

	phone_book := make(map[string]string) //made empty array of contact object - phone book

	var input int = 0
	var running = true
	var name string
	var phone string
	var del string // number that is going to be deleted

	for running {
		fmt.Print("Welcome to phone adress book app\n")
		fmt.Println("[1]- Add contact")
		fmt.Println("[2]- Remove contact")
		fmt.Println("[3]- View contacts")
		fmt.Println("[0]- Exit\n---(Navigate by typing index)---")

		fmt.Scanln(&input)
		switch {
		case input == 0:
			running = false // break statement
		case input == 1:
			// add contact
			fmt.Println("Enter name and phone number: ")
			fmt.Scanln(&name, &phone)
			phone_book[name] = phone // for key name assign value of variable phone
			row := db.QueryRowContext(context.Background(), "INSERT INTO contact (name, phone) VALUES(?, ?)", name, phone)
			if err := row.Err(); err != nil {
				fmt.Println("db.QueryRowContext", err)
				return
			}
		case input == 2:
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
		case input == 3:
			if len(phone_book) == 0 {
				fmt.Println("No contacts")
				break
			} else {
				i := 1
				for key, value := range phone_book {
					fmt.Printf("%v. %v: %v\n", i, key, value)
					i++
				}
			}
		default:
			fmt.Println("Invalid input")
			continue
		}
	}
}
