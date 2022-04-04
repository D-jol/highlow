package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var life int = 4 // make - player chooses difficulty, more or less lifes
	fmt.Print("you start with ", life, " lifes\n")

	s1 := rand.NewSource(time.Now().UnixNano()) //podesavanje parametra za random broj
	r1 := rand.New(s1)                          //

	var number int
	number = r1.Intn(10) //generisanje nasumicnog broja u rasponu od 0 do 10
	fmt.Println("number is ", number)
	var guess int

	for {
		fmt.Print("guess a number: ")
		fmt.Scanln(&guess) // uneta vrednost se prosledjuje quess varijabli

		if guess == number {
			fmt.Printf("Good job!You won with %v lifes left! \nExiting app...", life)
			time.Sleep(4 * time.Second)
			break
		} else if guess > number {
			life -= 1
			if life != 0 {
				fmt.Println("Try lower.")
			}
			//fmt.Println("Try lower.")
		} else if guess < number {
			life -= 1
			if life != 0 {
				fmt.Println("Try higher.")
			}
			//fmt.Println("Try higher.")
		}
		fmt.Printf("%v lifes left\n", life)

		if life == 0 {
			fmt.Println("Game over")
			fmt.Println("Exiting. . .")
			time.Sleep(3 * time.Second)
			break
		}
	}
}
