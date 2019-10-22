package main

import ("fmt"
		"os"
		"bufio"
)

func main() {

	planets := map[string]string{
		"Jupiter": "Jupiter is perpetually covered with clouds composed of ammonia crystals.",
		"Venus":   "Venus lies within Earth's orbit, and so doesn't venture far from the Sun, setting in the west after dusk and rising in the east before dawn.",
		"Mars": "Mars is also referred to as the 'Red Planet' because of the effect of the iron oxide prevalent on Mars' surface.",
		"Earth": "Earth is inhabited by humans, most of whom suffer from Wealth Accumulation Disorder.",
		"Neptune": "Neptune is cloudly and blue because the red light is absorbed  by the atmospheric methane.",
	}

	fmt.Println("Welcome to the Solar System!\nWhat is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("Nice to meet you ", name, ".\nMy name is Eliza, I'm an old friend of Alexa.")

	fmt.Println("Let's go on an adventure!\nShall I randomly choose a planet for you to visit? Y/N")
	var yorn string
	fmt.Scanln(&yorn)

	if yorn == "Y" {
		fmt.Println("I've selected the planet Jupiter, buckle up Buttercup!")
	} else {
		fmt.Println("What planet would you like to visit?")
		r := bufio.NewReader(os.Stdin)
		planet, _ := r.ReadString('\n')
		for k, v := range planets { 
			if k == planet{
			fmt.Printf(v)
			}
		}
	}

}