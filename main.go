package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pockedex > ")

		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		input := strings.Fields(strings.ToLower(scanner.Text()))
		fmt.Printf("Your command was: %v \n", input[0])
	}
}
