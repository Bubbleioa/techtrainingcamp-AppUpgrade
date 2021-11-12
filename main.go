package main

import (
	"fmt"
	"net/http"
	"os"
	"techtrainingcamp-AppUpgrade/database"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	lst, _ := database.QueryAllRules()
	for index, _ := range *lst {
		fmt.Println((*lst)[index]["id"])
		database.RedisTouchRule((*lst)[index]["id"])
	}
	r := gin.Default()
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	customizeouter(r)
	go srv.ListenAndServe()

	r2 := gin.Default()
	if os.Getenv("IS_DOCKER") == "1" {
		r2.LoadHTMLGlob("/root/public/index.html")
		database.MysqlCreateTable()
	} else {
		r2.LoadHTMLGlob("./public/index.html")
	}
	srv2 := &http.Server{
		Addr:         ":11451",
		Handler:      r2,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	adminRouter(r2)
	srv2.ListenAndServe()
	//r2.Run(":11451")
}
