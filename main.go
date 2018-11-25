package main

import (
	"fmt"

	"github.com/doovemax/log_trace/config"
)

func main() {
	config.Init()

	// fmt.Println(config.ConfigPath)
	host := config.Conf.Hosts["web1"]
	fmt.Printf("%+v\n", host)
	err := host.ReadFile("/var/log/nginx/access.log")
	fmt.Println(err)
}
