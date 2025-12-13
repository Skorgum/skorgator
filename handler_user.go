package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	// Implementation for login command
	if len(cmd.args) < 1 {
		return fmt.Errorf("username is required")
	}

	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}
