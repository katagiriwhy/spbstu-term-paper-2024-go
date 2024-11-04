package main

import (
	"fmt"
	"log"

	"github.com/katagiriwhy/database/config"
	db2 "github.com/katagiriwhy/database/db"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := db2.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
}
