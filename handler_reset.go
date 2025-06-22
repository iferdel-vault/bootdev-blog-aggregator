package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting all users from table: %v", err)
	}

	fmt.Println("users table successfully reset")
	return nil
}
