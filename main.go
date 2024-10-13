package main

import (
	"fmt"
	"os"
	"strconv"
)

var strength = 0
var health = 0
var stamina = 0
var intelligence = 0
var last_used = 0

func main() {
	print_welcome()
	print_menu()
	for {

		var selection string
		fmt.Scanln(&selection)

		selection_int, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Println("Invalid selection. Please enter a number between 1 and 6.")
			continue
		}

		if last_used == selection_int {
			fmt.Println("You can not raise the same stat twice in a row.")
			continue
		}

		switch selection_int {
		case 1:
			strength_up()
		case 2:
			health_up()
		case 3:
			stamina_up()
		case 4:
			intelligence_up()
		case 5:
			check_stats()
		case 6:
			fmt.Println("Goodbye! Thanks for playing!")
			os.Exit(0)
		default:
			fmt.Println("Invalid selection. Please choose a number between 1 and 6.")
		}

		if strength >= 10 || health >= 10 || stamina >= 10 || intelligence >= 10 {
			fmt.Println("Congratulations! You have reached the maximum stat of 10!")
			os.Exit(0)
		}

	}
}
