package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(c *gin.Context) {
		s:=time.Now()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
		//log latency, response code, path
		logger.Info("incoming request", zap.String("path", c.Request.URL.Path), zap.Int("status", c.Writer.Status()),zap.Duration("elapsed",time.Now().Sub(s)))
		//c.Next()
	}, func(c *gin.Context) {
		c.Set("requestId",rand.Int())
		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
