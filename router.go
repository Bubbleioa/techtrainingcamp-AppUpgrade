package main

import (
	"techtrainingcamp-AppUpgrade/service"

	"github.com/gin-gonic/gin"
)

func customizeouter(r *gin.Engine) {
	r.GET("/ping", service.Pong)
	r.GET("/judge", service.Hit)
}
