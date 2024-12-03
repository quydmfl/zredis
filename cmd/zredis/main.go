package main

import (
	"fmt"

	"github.com/quydmfl/zredis/internal/config"
)

func main() {
	conf, err := config.LoadConfig("./config")

	if err != nil {
		panic(err)
	}

	fmt.Printf("RunID: %s\n", conf.RunID)
	fmt.Printf("Bind: %s\n", conf.Bind)
	fmt.Printf("Port: %d\n", conf.Port)
	fmt.Printf("Timeout: %s\n", conf.Timeout)
}