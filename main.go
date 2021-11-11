package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	customizeouter(r)
	go r.Run()

	r2 := gin.Default()
	r2.LoadHTMLGlob("/root/public/index.html")
	adminRouter(r2)
	r2.Run(":11451")
}
