package service

import (
	"techtrainingcamp-AppUpgrade/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func Hit(c *gin.Context) {

	var respUrl string
	appVersion := c.Query("appVersion")
	userDID := c.Query("userDID")
	rules := model.GetAllRules()

	for index := 0; index < len(*rules); index++ {
		if cast.ToInt(userDID) < (*rules)[index].MinUserDID || cast.ToInt(userDID) > (*rules)[index].MaxUserDID {
			continue
		}

		if cast.ToInt(appVersion) < (*rules)[index].MinVersion || cast.ToInt(appVersion) > (*rules)[index].MaxVersion {
			continue
		}

		respUrl = (*rules)[index].GrayLink
		break
	}
	c.JSON(200, gin.H{"downloadUrl": respUrl})
}
