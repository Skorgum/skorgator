package main

import "context"

func handlerReset(s *state, cmd command) error {
	// Delete all users from the database
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}
