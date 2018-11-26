package module

import (
	"fmt"
	"strings"

	"github.com/doovemax/log_trace/config"
	"golang.org/x/net/websocket"
)

func init() {

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
	logTrans := make(chan []byte, 10000)
	for ser := range config.Conf.LogFile[serviceName] {
		fmt.Println(ser)
	}
}
