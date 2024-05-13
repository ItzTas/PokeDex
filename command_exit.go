package main

import (
	"fmt"
	"os"
	"slices"
)

func commandExit(c *config, s string) error {
	if s == "" {
		executeCommandExit()
		return nil
	}
	if slices.Contains(paramsHelp, s) {
		commandExitHelp()
	}
	return fmt.Errorf("unkown second parameter: %s", s)
}

func executeCommandExit() {
	os.Exit(0)
}

func commandExitHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["exit"].name, GetCommands()["exit"].description)
}