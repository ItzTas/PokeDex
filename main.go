package main

import (
	"fmt"
)

func main() { 
	err := GetCommands()["map"].callback()
	if err != nil{
		fmt.Println("error")
		fmt.Print(err)
	}
	startRepl()
}
