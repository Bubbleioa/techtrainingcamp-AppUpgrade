package service

import (
	_ "database/sql"
	"log"
	"strconv"
	"strings"
	"techtrainingcamp-AppUpgrade/database"
	"techtrainingcamp-AppUpgrade/tools"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// @title    qrnVersionComp
// @description   用于比较两个字符串类型的版本号
// @param     s1        string         "第一个版本号"
// @param     s2        string         "第二个版本号"
// @param     flag        bool         "true判断s1是否大于等于s2，false则判断s1是否小于等于s2"
// @return    无        bool         "判断结果"
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

// Judge 是一个路由函数，根据现有规则对用户发送的更新查询作出响应
func Judge(c *gin.Context) {

	var respUrl string
	var respUpdateVersionCode string
	var respMd5 string
	var respTitle string
	var respUpdateTips string
	var respRuleId string

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

	respRuleId, respUrl, respUpdateVersionCode, respMd5, respTitle, respUpdateTips = judgeLogic(idList, deviceId, aid, devicePlatform, cpuArch, channel, osApi, updateVersionCode)
	c.JSON(200, gin.H{"ruleid": respRuleId, "downloadUrl": respUrl, "UpdateVersionCode": respUpdateVersionCode,
		"Md5": respMd5, "Title": respTitle, "UpdateTips": respUpdateTips})
}

func judgeLogic(idList *[]string, deviceId string, aid string, devicePlatform string, cpuArch string, channel string, osApi string, updateVersionCode string) (string, string, string, string, string, string) {
	var respUrl string
	var respUpdateVersionCode string
	var respMd5 string
	var respTitle string
	var respUpdateTips string
	var respRuleId string
	for index := 0; index < len(*idList); index++ {
		ruleid := (*idList)[index]

		qObj := database.RuleObj{}
		isEnabled, _ := qObj.GetRuleAtt(ruleid, "enabled")
		if !cast.ToBool(isEnabled) {
			continue
		}
		ruleAid, _ := qObj.GetRuleAtt(ruleid, "aid")
		rulePlatform, _ := qObj.GetRuleAtt(ruleid, "platform")
		ruleCpuArch, _ := qObj.GetRuleAtt(ruleid, "cpu_arch")
		ruleChannel, _ := qObj.GetRuleAtt(ruleid, "channel")
		isDeviceIDValue, _ := qObj.CheckDeviceIDInWhiteList(ruleid, deviceId)
		ruleMinOsApi, _ := qObj.GetRuleAtt(ruleid, "min_os_api")
		ruleMaxOsApi, _ := qObj.GetRuleAtt(ruleid, "max_os_api")
		ruleMinUpdateVersionCode, _ := qObj.GetRuleAtt(ruleid, "min_update_version_code")
		ruleMaxUpdateVersionCode, _ := qObj.GetRuleAtt(ruleid, "max_update_version_code")
		if strings.Compare(aid, ruleAid) == 0 &&
			strings.Compare(devicePlatform, rulePlatform) == 0 &&
			strings.Compare(channel, ruleChannel) == 0 {
			if isDeviceIDValue {
				respRuleId = ruleid
				respUrl, _ = qObj.GetRuleAtt(ruleid, "download_url")
				respUpdateVersionCode, _ = qObj.GetRuleAtt(ruleid, "update_version_code")
				respMd5, _ = qObj.GetRuleAtt(ruleid, "md5")
				respTitle, _ = qObj.GetRuleAtt(ruleid, "title")
				respUpdateTips, _ = qObj.GetRuleAtt(ruleid, "update_tips")
				break
			} else {
				if cast.ToInt(osApi) >= cast.ToInt(ruleMinOsApi) &&
					cast.ToInt(osApi) <= cast.ToInt(ruleMaxOsApi) &&
					tools.VersionCmp(updateVersionCode, ruleMinUpdateVersionCode) != -1 &&
					tools.VersionCmp(updateVersionCode, ruleMaxUpdateVersionCode) != 1 &&
					strings.Compare(cpuArch, ruleCpuArch) == 0 {
					respRuleId = ruleid
					respUrl, _ = qObj.GetRuleAtt(ruleid, "download_url")
					respUpdateVersionCode, _ = qObj.GetRuleAtt(ruleid, "update_version_code")
					respMd5, _ = qObj.GetRuleAtt(ruleid, "md5")
					respTitle, _ = qObj.GetRuleAtt(ruleid, "title")
					respUpdateTips, _ = qObj.GetRuleAtt(ruleid, "update_tips")
					break
				}
			}
		}
	}
	// for index := 0; index < len(*idList); index++ {
	// 	ruleid := (*idList)[index]
	// 	isEnabled, _ := database.GetRuleAtt(ruleid, "enabled")
	// 	if !cast.ToBool(isEnabled) {
	// 		continue
	// 	}
	// 	ruleAid, _ := database.GetRuleAtt(ruleid, "aid")
	// 	rulePlatform, _ := database.GetRuleAtt(ruleid, "platform")
	// 	ruleCpuArch, _ := database.GetRuleAtt(ruleid, "cpu_arch")
	// 	ruleChannel, _ := database.GetRuleAtt(ruleid, "channel")
	// 	isDeviceIDValue, _ := database.CheckDeviceIDInWhiteList(ruleid, deviceId)
	// 	ruleMinOsApi, _ := database.GetRuleAtt(ruleid, "min_os_api")
	// 	ruleMaxOsApi, _ := database.GetRuleAtt(ruleid, "max_os_api")
	// 	ruleMinUpdateVersionCode, _ := database.GetRuleAtt(ruleid, "min_update_version_code")
	// 	ruleMaxUpdateVersionCode, _ := database.GetRuleAtt(ruleid, "max_update_version_code")
	// 	if strings.Compare(aid, ruleAid) == 0 &&
	// 		strings.Compare(devicePlatform, rulePlatform) == 0 &&
	// 		strings.Compare(cpuArch, ruleCpuArch) == 0 &&
	// 		strings.Compare(channel, ruleChannel) == 0 &&
	// 		isDeviceIDValue &&
	// 		cast.ToInt(osApi) >= cast.ToInt(ruleMinOsApi) &&
	// 		cast.ToInt(osApi) <= cast.ToInt(ruleMaxOsApi) &&
	// 		tools.VersionCmp(updateVersionCode, ruleMinUpdateVersionCode) != -1 &&
	// 		tools.VersionCmp(updateVersionCode, ruleMaxUpdateVersionCode) != 1 {
	// 		respUrl, _ = database.GetRuleAtt(ruleid, "download_url")
	// 		respUpdateVersionCode, _ = database.GetRuleAtt(ruleid, "update_version_code")
	// 		respMd5, _ = database.GetRuleAtt(ruleid, "md5")
	// 		respTitle, _ = database.GetRuleAtt(ruleid, "title")
	// 		respUpdateTips, _ = database.GetRuleAtt(ruleid, "update_tips")
	// 		respRuleId, _ = database.GetRuleAtt(ruleid, "id")
	// 		break
	// 	}

	// }
	return respRuleId, respUrl, respUpdateVersionCode, respMd5, respTitle, respUpdateTips
}

func Count(c *gin.Context) {
	ruleId := c.Query("ruleid")
	isDownload := c.Query("download")
	var err error
	if cast.ToInt(isDownload) == 1 {
		err = database.UpdateUserDownloadStatus(ruleId, true)
	} else {
		err = database.UpdateUserDownloadStatus(ruleId, false)
	}
	tools.LogMsg(err)
	c.JSON(200, gin.H{})
}
