package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
