package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// database.RedisInitClient()
	// database.OpenMysql()
	// defer database.RedisClose()
	// defer database.CloseMysql()
	// lst, _ := database.QueryAllRules()
	// for index, _ := range *lst {
	// 	fmt.Println((*lst)[index]["id"])
	// 	database.RedisTouchRule((*lst)[index]["id"])
	// }
	// f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	r := gin.Default()
	customizeouter(r)
	r.Run(":8080")

	// r2 := gin.Default()
	// if os.Getenv("IS_DOCKER") == "1" {
	// 	r2.LoadHTMLGlob("/root/public/index.html")
	// 	database.MysqlCreateTable()
	// } else {
	// 	r2.LoadHTMLGlob("./public/index.html")
	// }
	// adminRouter(r2)
	// go r2.Run(":11451")
	// time.Sleep(20 * time.Second)
	// panic("NO!!!!!")
	// r2.Run(":11451")
}
