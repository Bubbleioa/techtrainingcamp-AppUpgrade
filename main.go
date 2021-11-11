package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	customizeouter(r)
	go r.Run()

	r2 := gin.Default()
	if os.Getenv("IS_DOCKER") == "1" {
		r2.LoadHTMLGlob("/root/public/index.html")
	} else {
		r2.LoadHTMLGlob("./public/index.html")
	}
	adminRouter(r2)
	r2.Run(":11451")
}
