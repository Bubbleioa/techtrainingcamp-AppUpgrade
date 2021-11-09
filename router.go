package main

import (
	"github.com/gin-gonic/gin"
	"techtrainingcamp-AppUpgrade/service"
)

func customizeouter(r *gin.Engine) {
	//r.GET("/ping", service.Pong)
	//r.GET("/judge1", service.Hit)
	//r.GET("/judge2", service.HitSQL)
	r.GET("/index",service.test)
}
