package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrade = websocket.Upgrader{
	HandshakeTimeout:  10 * time.Second,
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("---check origin---")
		return true
	},
	EnableCompression: false,
}
func ping(c *gin.Context) {
	fmt.Println("----in ping-----")
	// 升级get 请求到websocket 请求
	var err error
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("---err-----", err)
		return 
	}
	fmt.Println("------111-----")
	defer ws.Close()
	for {
		// 读取ws数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("---err1----", err)
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		// 写入数据到ws
		ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("-----err2-----", err)
			break
		}
		
	}
	
}
func main() {
	engine := gin.Default()
	engine.GET("/ping", ping)
	
	engine.Run(":2023")
}
