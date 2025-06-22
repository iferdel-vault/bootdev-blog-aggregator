package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/iferdel-vault/bootdev-blog-aggregator/internal/config"
	"github.com/iferdel-vault/bootdev-blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error with reading config file: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error opening connection to DBUrl")
	}
	dbQueries := database.New(db)

	var s state
	s.cfg = &cfg
	s.db = dbQueries

	cmds := commands{
		cmdList: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
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
