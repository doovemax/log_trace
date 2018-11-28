package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"

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
	// logrus.Infoln("Server runing Success at", config.Conf.ServerConfig["BindAddr"], config.Conf.ServerConfig["ListenPort"], "!")
	http.Handle("/*", websocket.Handler(LogSh))

	if err := http.ListenAndServe(":"+config.Conf.ServerConfig.ListenPort, nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func LogSh(ws *websocket.Conn) {
	url := ws.Request().URL.Path
	urlSlice := strings.Split(url, "/")
	if len(urlSlice) != 3 {
		ws.Write([]byte("url error"))
		return
	}
	// projectName := urlSlice[1]
	serviceName := urlSlice[2]
	// logTrans := make(chan []byte, 10000)
	for ser := range config.Conf.LogFile[serviceName] {
		fmt.Println(ser)

	}
}
