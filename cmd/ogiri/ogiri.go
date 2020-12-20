package main

import (
	"fmt"
	"github.com/hatobus/o-giri/config"
	"os"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] faild to load config. err: %v", err)
		os.Exit(1)
	}
	fmt.Println(conf)
}
