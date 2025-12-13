package main

import (
	"fmt"
	"log"

	"github.com/skorgum/skorgator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("skorgum")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
}
