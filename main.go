package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/iferdel-vault/bootdev-blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

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

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error with reading config file: %v", err)
	}
	err = cfg.SetUser("iferdel")
	if err != nil {
		log.Fatalf("error with setting user in config file: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error with reading config file: %v", err)
	}
	fmt.Println(cfg)
}
