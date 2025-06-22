package main

import (
	"errors"
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	cmdList map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	f, ok := cmds.cmdList[cmd.Name]
	if !ok {
		return errors.New("command does not exists in available commands")
	}
	return f(s, cmd)
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	_, ok := cmds.cmdList[name]
	if ok {
		fmt.Printf("command %s already exists\n", name)
		return
	}
	cmds.cmdList[name] = f
}
