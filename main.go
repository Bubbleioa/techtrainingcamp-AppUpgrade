package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	customizeouter(r)
	go r.Run()

	r2 := gin.Default()
	adminRouter(r2)
	r2.Run(":11451")
}
