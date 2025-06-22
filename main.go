package main

import (
	"log"
	"os"

	"github.com/iferdel-vault/bootdev-blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error with reading config file: %v", err)
	}

	var s state
	s.cfg = &cfg

	cmds := commands{
		cmdList: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	stdin := os.Args

	if len(stdin) < 2 {
		log.Fatalf("not enough arguments passed, exiting...")
	}

	args := stdin[1:]
	cmd := command{
		Name: args[0],
		Args: args[1:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("error with running command: %v", err)
	}

}
