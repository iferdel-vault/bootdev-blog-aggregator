package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	username := cmd.args[0]

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("user %s has been set\n", username)

	return nil
}
