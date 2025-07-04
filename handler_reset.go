package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete users: %w", err)
	}

	fmt.Println("users table successfully reset")
	return nil
}
