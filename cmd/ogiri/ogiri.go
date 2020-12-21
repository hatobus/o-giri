package main

import (
	"fmt"
	"os"

	"github.com/hatobus/o-giri/config"
	"github.com/hatobus/o-giri/infrastructure/database"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] faild to load config. err: %v", err)
		os.Exit(1)
	}

	db, err := database.Connect(conf.MySQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] faild to connect database. err: %v", err)
		os.Exit(1)
	}

	fmt.Println(db.Stats())
}
