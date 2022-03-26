package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var life int = 3 // make - player chooses difficulty, more or less lifes
	fmt.Print("you start with ", life, " lifes\n")

	s1 := rand.NewSource(time.Now().UnixNano()) //podesavanje parametra za random broj
	r1 := rand.New(s1)                          //

	var number int
	number = r1.Intn(10) //generisanje nasumicnog broja u rasponu od 0 do 10
	fmt.Println("number is ", number)
	var guess int

	for life > 0 {
		fmt.Print("guess a number: ")
		fmt.Scanln(&guess) // uneta vrednost se prosledjuje quess varijabli

		if guess == number {
			fmt.Printf("Good job!You won with %v lifes left! \nExiting app...", life)
			time.Sleep(4 * time.Second)
			life -= life
		} else if guess > number {
			life -= 1
			fmt.Println("Try lower.")
		} else if guess < number {
			life -= 1
			fmt.Println("Try higher.")
		}

	}
	if life == 0 {
		fmt.Print("No more lifes.Game over")
	}
}
