package main

import (

	"techtrainingcamp-AppUpgrade/admin"

	"github.com/gin-gonic/gin"
	"techtrainingcamp-AppUpgrade/service"
)

func customizeouter(r *gin.Engine) {

	r.GET("/index",service.test)

	r.GET("/ping", service.Pong)
	r.GET("/judge1", service.Hit)
	r.GET("/judge2", service.HitSQL)
	r.GET("/judge", service.Judge)
	r.GET("/count", service.Count)

}

func adminRouter(r *gin.Engine) {
	r.GET("/query_all_rules", admin.QueryAllRules)
	r.GET("/query_rule", admin.QueryRule)
	r.POST("/update_rule", admin.UpdateRule)
	r.POST("/create_rule", admin.CreateRule)
	r.GET("/delete_rule", admin.DeleteRule)
}
