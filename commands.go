package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmdList map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {

	f, ok := cmds.cmdList[cmd.name]
	if !ok {
		return errors.New("command does not exists in avilable commands")
	}
	err := f(s, cmd)
	if err != nil {
		return fmt.Errorf("error running command %s: %w", cmd.name, err)
	}

	return nil
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	_, ok := cmds.cmdList[name]
	if ok {
		fmt.Printf("command %s already exists\n", name)
		return
	}
	cmds.cmdList[name] = f
}
