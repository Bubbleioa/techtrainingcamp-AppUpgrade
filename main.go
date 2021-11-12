package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime/pprof"
	"techtrainingcamp-AppUpgrade/database"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	database.RedisInitClient()
	database.OpenMysql()
	defer database.RedisClose()
	defer database.CloseMysql()
	lst, _ := database.QueryAllRules()
	for index, _ := range *lst {
		fmt.Println((*lst)[index]["id"])
		database.RedisTouchRule((*lst)[index]["id"])
	}
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
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
	go srv2.ListenAndServe()
	time.Sleep(20 * time.Second)
	panic("NO!!!!!")
	//r2.Run(":11451")
}
