package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()
	engine.LoadHTMLFiles("/usr/local/works/go/src/go-test/module/test-websocket/client/templates/index.tmpl")
	engine.GET("/admin/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"body": "你好body",
		})
		return
	})
	engine.Run(":2024")
}
