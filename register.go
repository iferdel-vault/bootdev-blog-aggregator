package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iferdel-vault/bootdev-blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("the register command expects a single argument, the username.")
	}
	username := cmd.Args[0]

	// user, err := s.db.GetUser(context.Background(), username)
	// if err != nil {
	// 	return err
	// }
	// if user {
	// 	os.Exit(1)
	// }

	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("user %s has been created & set\n", username)
	return nil
}
