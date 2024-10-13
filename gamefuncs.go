package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

func strength_up() {
	clear_screen()
	strength = increase_stat(strength)
	health = decrease_stat(health)
	stamina = decrease_stat(stamina)
	intelligence = decrease_stat(intelligence)
	last_used = 1
	fmt.Println("You raised your strength!")
	check_stats()
	print_menu()
}

func health_up() {
	clear_screen()
	health = increase_stat(health)
	strength = decrease_stat(strength)
	stamina = decrease_stat(stamina)
	intelligence = decrease_stat(intelligence)
	last_used = 2
	fmt.Println("You raised your health!")
	check_stats()
	print_menu()
}

func stamina_up() {
	clear_screen()
	stamina = increase_stat(stamina)
	strength = decrease_stat(strength)
	health = decrease_stat(health)
	intelligence = decrease_stat(intelligence)
	last_used = 3
	fmt.Println("You raised your stamina!")
	check_stats()
	print_menu()
}

func intelligence_up() {
	clear_screen()
	intelligence = increase_stat(intelligence)
	strength = decrease_stat(strength)
	stamina = decrease_stat(stamina)
	health = decrease_stat(health)
	last_used = 4
	fmt.Println("You raised your intelligence!")
	check_stats()
	print_menu()
}

func increase_stat(stat int) int {
	stat += rand.Intn(5) + 1
	if stat > 10 {
		stat = 10
	}	
	return stat
}
func decrease_stat(stat int) int {
	stat -= rand.Intn(3) + 1
	if stat <= 0 {
		stat = 0
	}
	return stat
}

func check_stats() {
	fmt.Printf("Your current stats are: \nStrength: %d\nHealth: %d\nStamina: %d\nIntelligence: %d\n", strength, health, stamina, intelligence)
}

func print_menu() {
	fmt.Println("Please enter the number of the stat you want to raise:")
	fmt.Println("1. Strength")
	fmt.Println("2. Health")
	fmt.Println("3. Stamina")
	fmt.Println("4. Intelligence")
	fmt.Println("5. To see your current stats")
	fmt.Println("6. To quit")
}

func print_welcome() {
	fmt.Println("Welcome to the Stat Game!")
	fmt.Println("The goal is to get at least one stat to 10.")
	fmt.Println("You cannot raise one stat without lowering the others.")
	fmt.Println("You cannot raise the same stat twice in a row.")
}

func clear_screen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}