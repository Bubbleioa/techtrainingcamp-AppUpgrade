// 路由的回调函数
package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"techtrainingcamp-AppUpgrade/model"
	"database/sql"
	"github.com/spf13/cast"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// 逐条判断规则
func Hit(c *gin.Context) {

	rules := model.GetRulesFromFile()

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
func versionComp(s1 string, s2 string, flag bool) bool {
	arr1 := strings.Split(s1, ".")
	arr2 := strings.Split(s2, ".")
	for index := 0; index < len(arr1); index++ {
		intTemp, err := strconv.Atoi(arr1[index])
		if err != nil {
			log.Panic(err.Error())
		}
		arr1[index] = strconv.Itoa(intTemp)
	}
	for index := 0; index < len(arr2); index++ {
		intTemp, err := strconv.Atoi(arr2[index])
		if err != nil {
			log.Panic(err.Error())
		}
		arr2[index] = strconv.Itoa(intTemp)
	}
	arrLen := len(arr2)
	if len(arr1) < len(arr2) {
		arrLen = len(arr1)
	}
	if flag {
		for index := 0; index < arrLen; index++ {
			if strings.Compare(arr1[index], arr2[index]) < 0 {
				return false
			}
		}
		return true
	} else {
		for index := 0; index < arrLen; index++ {
			if strings.Compare(arr1[index], arr2[index]) > 0 {
				return false
			}
		}
		return true
	}
}

//判断规则+数据库
func HitSQL(c *gin.Context) {

	var respUrl string
	var respUpdateVersionCode string
	var respMd5 string
	var respTitle string
	var respUpdateTips string
	devicePlatform := c.Query("devicePlatform")
	deviceId := c.Query("deviceId")
	osApi := c.Query("osApi")
	channel := c.Query("channel")
	updateVersionCode := c.Query("updateVersionCode")
	cpuArch := c.Query("cpuArch")
	rules := model.GetAllRules()
	for index := 0; index < len(*rules); index++ {
		if strings.Compare(devicePlatform, (*rules)[index].Platform) == 0 &&
			strings.Compare(cpuArch, (*rules)[index].CpuArch) == 0 &&
			strings.Compare(channel, (*rules)[index].Channel) == 0 {
			constr := "root:ru19870528@tcp(127.0.0.1:3306)/ginsql"
			//打开连接
			db, err := sql.Open("mysql", constr) //返回mysql实例db
			if err != nil {
				log.Panic(err.Error())
				return
			}
			rows, err := db.Query("select device_id_list from device_id where device_id_list=" + deviceId)
			if err != nil {
				log.Panic(err.Error())
				return
			}
			var flag bool
			flag = false
			for rows.Next() {
				var id string
				err := rows.Scan(&id) //读取rows里面的数据分别赋值给结构体属性
				if err != nil {
					log.Panic(err.Error())
					return
				}
				if id == deviceId {
					flag = true
					break
				}
			}
			if flag && cast.ToInt(osApi) >= (*rules)[index].MinOsApi &&
				cast.ToInt(osApi) <= (*rules)[index].MaxOsApi &&
				versionComp(updateVersionCode, (*rules)[index].MinUpdateVersionCode, true) &&
				versionComp(updateVersionCode, (*rules)[index].MaxUpdateVersionCode, false) {
				respUrl = (*rules)[index].DownloadUrl
				respUpdateVersionCode = (*rules)[index].UpdateVersionCode
				respMd5 = (*rules)[index].Md5
				respTitle = (*rules)[index].Title
				respUpdateTips = (*rules)[index].UpdateTips
				break
			}
		}
	}
	c.JSON(200, gin.H{"downloadUrl": respUrl, "UpdateVersionCode": respUpdateVersionCode,
		"Md5": respMd5, "Title": respTitle, "UpdateTips": respUpdateTips})
}
