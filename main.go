package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() { 
	for {
		c := GetCommands()
		scanner := bufio.NewScanner(os.Stdin)	
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		
		if input == c["exit"].name { 
			err := c["exit"].callback()
			if err != nil {
                fmt.Println("Error: ", err)
            }
		}

		err := c[input].callback

		if err != nil {
			fmt.Println("Error") // Fix it
		}
	}
}
