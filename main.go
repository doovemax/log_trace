package main

import (
	"fmt"
	"net/http"

	"github.com/doovemax/log_trace/module"

	"golang.org/x/net/websocket"

	"github.com/sirupsen/logrus"

	"github.com/doovemax/log_trace/config"
)

func main() {
	config.Init()

	fmt.Println(config.ConfigPath)
	// host := config.Conf.Hosts["web1"]
	// fmt.Printf("%+v\n", host)
	// err := host.ReadFile("/var/log/nginx/access.log")
	// fmt.Println(err)
	// http.Handle("/echo", websocket.Handler(echoHandler))
	// http.Handle("/", http.FileServer(http.Dir(".")))
	mp := config.Conf.ServerConfig
	logrus.Info(mp["BindAddr"])
	logrus.Infoln("Server runing Success at", config.Conf.ServerConfig["BindAddr"], config.Conf.ServerConfig["ListenPort"], "!")
	http.HandleFunc("/*", websocket.Handler(module.LogSh))

	if err := http.ListenAndServe(":"+config.Conf.ServerConfig["ListenPort"], nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
