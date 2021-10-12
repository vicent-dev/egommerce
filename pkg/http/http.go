package http

import "github.com/gin-gonic/gin"

func InitServer() error {
	r := gin.Default()
	return r.Run()
}
