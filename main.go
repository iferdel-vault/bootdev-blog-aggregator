package main

import (
	"fmt"
	"log"

	"github.com/iferdel-vault/bootdev-blog-aggregator/internal/config"
)

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
