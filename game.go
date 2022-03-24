package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var life int = 6
	fmt.Print("you start with 6 lifes\n")

	s1 := rand.NewSource(time.Now().UnixNano()) //podesavanje parametra za random broj
	r1 := rand.New(s1)                          //

	var number int
	number = r1.Intn(10) //generisanje nasumicnog broja u rasponu od 0 do 10
	//fmt.Println("number is ", number)
	var quess int
	//fmt.Print(number)
	for life != 0 {
		fmt.Print("quess a number: ")
		fmt.Scanln(&quess) //uneta vrednost se prosledjuje quess varijabli
		if quess == number {
			fmt.Println("correct")
			fmt.Print("You won! Exiting app...")
			time.Sleep(4 * time.Second)
			break
		} else {
			life -= 1

			if life > 0 {
				fmt.Println("Life count: ", life)
			} else {
				fmt.Print("You lost !")
			}
		}

	}

}
