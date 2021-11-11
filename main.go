package main

import (
	"fmt"
	"os"
	"techtrainingcamp-AppUpgrade/database"

	"github.com/gin-gonic/gin"
)

func main() {
	lst, _ := database.QueryAllRules()
	for index, _ := range *lst {
		fmt.Println((*lst)[index]["id"])
		database.RedisTouchRule((*lst)[index]["id"])
	}
	r := gin.Default()

	customizeouter(r)
	go r.Run()

	r2 := gin.Default()
	if os.Getenv("IS_DOCKER") == "1" {
		r2.LoadHTMLGlob("/root/public/index.html")
		database.MysqlCreateTable()
	} else {
		r2.LoadHTMLGlob("./public/index.html")
	}
	adminRouter(r2)
	r2.Run(":11451")
}
