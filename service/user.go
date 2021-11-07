package service

import (
	_ "database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"strconv"
	"strings"
	"techtrainingcamp-AppUpgrade/database"
	_ "techtrainingcamp-AppUpgrade/model"
)

/*param
s1 第一个字符串
s2 第二个字符串
flag 若为true，则判断s1版本是否大于等于s2版本，否则判断s1版本是否小于等于s2版本

return
判断结果
*/
func qrnVersionComp(s1 string, s2 string, flag bool) bool {
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

func qrnGetIDList() (*[]string, error) {
	var IDList = []string{"1", "2", "3"}
	return &IDList, nil
}

func qrnGetRuleAttr(ruleid string, field string) (string, error) {
	if field == "min_update_version_code" {
		return "8.4.0", nil
	} else if field == "max_update_version_code" {
		return "8.8.8", nil
	} else if field == "min_os_api" {
		return "10", nil
	} else if field == "max_os_api" {
		return "20", nil
	} else if field == "platform" {
		return "Android", nil
	} else if field == "cpu_arch" {
		return "32", nil
	} else if field == "channel" {
		return "dsd", nil
	} else if field == "download_url" {
		return "https://baidu1.com", nil
	} else if field == "update_version_code" {
		return "4.1", nil
	} else if field == "md5" {
		return "dsafaf", nil
	} else if field == "title" {
		return "升级啦", nil
	} else if field == "update_tips" {
		return "快升级", nil
	} else if strings.Compare(field, "aid") == 0 {
		return "8", nil
	} else if field == "enabled" {
		return "true", nil
	}
	return "", nil
}

func qrnCheckDeviceIDInWhiteList(ruleid string, deviceid string) (bool, error) {
	return true, nil
}

func qrnUpdateUserDownloadStatus(ruleid string, status bool) error {
	return nil
}

func Judge(c *gin.Context) {

	var respUrl string
	var respUpdateVersionCode string
	var respMd5 string
	var respTitle string
	var respUpdateTips string

	devicePlatform := c.Query("device_platform")
	deviceId := c.Query("device_id")
	osApi := c.Query("os_api")
	channel := c.Query("channel")
	updateVersionCode := c.Query("update_version_code")
	cpuArch := c.Query("cpu_arch")
	aid := c.Query("aid")
	idList, err := database.GetIDList()
	if err != nil {
		log.Panic(err.Error())
		return
	}

	for index := 0; index < len(*idList); index++ {
		ruleid := (*idList)[index]
		isEnabled, _ := database.GetRuleAtt(ruleid, "enabled")
		if !cast.ToBool(isEnabled) {
			continue
		}
		ruleAid, _ := database.GetRuleAtt(ruleid, "aid")
		rulePlatform, _ := database.GetRuleAtt(ruleid, "platform")
		ruleCpuArch, _ := database.GetRuleAtt(ruleid, "cpu_arch")
		ruleChannel, _ := database.GetRuleAtt(ruleid, "channel")
		isDeviceIDValue, _ := database.CheckDeviceIDInWhiteList(ruleid, deviceId)
		ruleMinOsApi, _ := database.GetRuleAtt(ruleid, "min_os_api")
		ruleMaxOsApi, _ := database.GetRuleAtt(ruleid, "max_os_api")
		ruleMinUpdateVersionCode, _ := database.GetRuleAtt(ruleid, "min_update_version_code")
		ruleMaxUpdateVersionCode, _ := database.GetRuleAtt(ruleid, "max_update_version_code")
		if strings.Compare(aid, ruleAid) == 0 &&
			strings.Compare(devicePlatform, rulePlatform) == 0 &&
			strings.Compare(cpuArch, ruleCpuArch) == 0 &&
			strings.Compare(channel, ruleChannel) == 0 &&
			isDeviceIDValue &&
			cast.ToInt(osApi) >= cast.ToInt(ruleMinOsApi) &&
			cast.ToInt(osApi) <= cast.ToInt(ruleMaxOsApi) &&
			qrnVersionComp(updateVersionCode, ruleMinUpdateVersionCode, true) &&
			qrnVersionComp(updateVersionCode, ruleMaxUpdateVersionCode, false) {
			respUrl, _ = database.GetRuleAtt(ruleid, "download_url")
			respUpdateVersionCode, _ = database.GetRuleAtt(ruleid, "update_version_code")
			respMd5, _ = database.GetRuleAtt(ruleid, "md5")
			respTitle, _ = database.GetRuleAtt(ruleid, "title")
			respUpdateTips, _ = database.GetRuleAtt(ruleid, "update_tips")
			break
		}

	}
	c.JSON(200, gin.H{"downloadUrl": respUrl, "UpdateVersionCode": respUpdateVersionCode,
		"Md5": respMd5, "Title": respTitle, "UpdateTips": respUpdateTips})
}

func Count(c *gin.Context) {
	ruleId := c.Query("ruleid")
	isDownload := c.Query("download")
	if cast.ToInt(isDownload) == 1 {
		_ = database.UpdateUserDownloadStatus(ruleId, true)
	} else {
		_ = database.UpdateUserDownloadStatus(ruleId, false)
	}
	c.JSON(200, gin.H{})
}
