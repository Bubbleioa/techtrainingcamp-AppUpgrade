package main

import (
	"techtrainingcamp-AppUpgrade/admin"
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

func adminRouter(r *gin.Engine) {
	r.GET("/get_all_rules", admin.QueryAllRules)
	r.GET("/get_rule", admin.QueryRule)
	r.POST("/update_rule", admin.UpdateRule)
	r.POST("/create_rule", admin.CreateRule)
	r.GET("/delete_rule", admin.DeleteRule)
}
