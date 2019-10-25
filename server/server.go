package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 获得当前gin框架引擎，对外主要是router
	engine := gin.Default()

	//  完成各项功能的router注册

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	// listen and serve on 0.0.0.0:8080
}
