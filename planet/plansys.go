package main

import ("fmt"
		"os"
		"bufio"
)
func main() {
	fmt.Println("Welcome to the Solar System!\nWhat is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Nice to meet you ", name, ".\nMy name is Eliza, I'm an old friend of Alexa.")
	fmt.Println("Let's go on an adventure!\nShall I randomly choose a planet for you to visit? Y/N")
	var yorn string 
	fmt.Scanline(&yorn)
	if yorn == 'Y' {
		fmt.Println("I've selected the planet Jupiter, buckle up Buttercup!")
	} else {
		fmt.Println("What planet would you like to visit?")
	}
}