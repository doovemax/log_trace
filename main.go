package main

import (
	"fmt"

	"github.com/doovemax/log_trace/config"
)

func main() {
	config.Init()

	fmt.Println(config.ConfigPath)
}
