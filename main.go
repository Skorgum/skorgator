package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skorgum/skorgator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{cfg: &cfg}

	commands := &commands{registeredCommands: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
	}
	cmd := command{os.Args[1], os.Args[2:]}
	err = commands.run(s, cmd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
