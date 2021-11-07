package admin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"techtrainingcamp-AppUpgrade/database"

	"github.com/gin-gonic/gin"
)

func QueryAllRules(c *gin.Context) {
	_, e := database.QueryAllRules()
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": e.Error()})
		return
	}

}

func QueryRule(c *gin.Context) {
	ruleid := c.Query("ruleid")
	_, _, e := database.QueryRuleByID(ruleid)
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": e.Error()})
	}
}

func UpdateRule(c *gin.Context) {
	str, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(str))
}

func CreateRule(c *gin.Context) {

}

func DeleteRule(c *gin.Context) {
	ruleid := c.Query("ruleid")
	e := database.DeleteRule(ruleid)
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": e.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}
