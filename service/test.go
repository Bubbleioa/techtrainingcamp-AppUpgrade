package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",nil)
}