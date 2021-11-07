package main

import (
	"techtrainingcamp-AppUpgrade/service"

	"github.com/gin-gonic/gin"
)

func customizeouter(r *gin.Engine) {
	r.GET("/ping", service.Pong)
	r.GET("/judge1", service.Hit)
	r.GET("/judge2", service.HitSQL)
	r.GET("/judge", service.Judge)
	r.GET("/count", service.Count)

}
