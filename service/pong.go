// 路由的回调函数
package service

import (
	"fmt"
	"techtrainingcamp-AppUpgrade/model"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// 逐条判断规则
func Hit(c *gin.Context) {

	rules := model.GetAllRules()

	for _, r := range *rules {
		cs := r.Hit
		flag := true
		for field, cnd := range cs {
			ok, err := cnd.SuccessStr(c.Query(field))
			if err != nil {
				fmt.Println(err)
				return
			}
			if !ok {
				fmt.Println(field, cnd, c.Query(field))
				flag = false
				break
			}
		}
		if flag {
			c.JSON(200, r.Res)
		}
	}
}
