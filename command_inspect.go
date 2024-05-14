package main

import (
	"fmt"
	"slices"
)

func commandInspect(cfg *config, s string) error {
	if slices.Contains(paramsHelp, s) {
		commandInspectHelp()
		return nil
	}
	if s != "" {
		return commandInspectexecute(cfg, s)
	}
	return fmt.Errorf("invalid second parameter: %s", s)
}

func commandInspectHelp() {
	fmt.Printf("\n %s: %s\n", GetCommands()["inspect"].name, GetCommands()["inspect"].description)
}

func commandInspectexecute(cfg *config, s string) error {
	p, ok := cfg.pokeGet(s)
	if !ok {
		return fmt.Errorf("you have not caught %s yet", s)
	}
	fmt.Println("")
	fmt.Println("Name:", p.Name)
	fmt.Println("Height:", p.Height)
	fmt.Println("weight:", p.Weight)
	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Println("  -" + stat.Stat.Name + ":", stat.BaseStat)
	} 
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Println("  -", t.Type.Name)
	}
	fmt.Println("")
	return nil
}