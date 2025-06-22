package main

import (
	"fmt"
	"log"

	"github.com/iferdel-vault/bootdev-blog-aggregator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
	err = cfg.SetUser("iferdel")
	if err != nil {
		log.Fatal(err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
