package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func VersionCmp(a string, b string) int {
	arr1 := strings.Split(a, ".")
	arr2 := strings.Split(b, ".")
	res1 := make([]int, len(arr1))
	res2 := make([]int, len(arr2))

	for index := 0; index < len(arr1); index++ {
		intTemp, err := strconv.Atoi(arr1[index])
		if err != nil {
			log.Panic(err.Error())
		}
		res1[index] = intTemp
	}
	for index := 0; index < len(arr2); index++ {
		intTemp, err := strconv.Atoi(arr2[index])
		if err != nil {
			log.Panic(err.Error())
		}
		res2[index] = intTemp
	}
	arrLen := len(arr2)
	if len(arr1) < len(arr2) {
		arrLen = len(arr1)
	}
	for index := 0; index < arrLen; index++ {
		if res1[index] < res2[index] {
			return -1
		} else if res1[index] > res2[index] {
			return 1
		}
	}
	return 0
}

func ConvertFullRuleToJSON(rule *map[string]string, devicelist *[]string) *string {
	if rule == nil && devicelist == nil {
		return nil
	}
	detail := make(map[string]interface{})
	for k, v := range *rule {
		detail[k] = v
	}
	detail["device_id_list"] = *devicelist
	mjson, _ := json.Marshal(detail)
	mString := string(mjson)
	//fmt.Println(mString)
	if mString == "" {
		return nil
	}
	return &mString
}

func ConvertSimplifiedRulesListToJson(rules *[]map[string]string) *string {
	if rules == nil {
		return nil
	}
	var ans string
	ans += "["
	for _, i := range *rules {
		mjson, _ := json.Marshal(i)
		mString := string(mjson)
		ans += mString + ","
	}
	ans = ans[0 : len(ans)-1]
	ans += "]"
	//fmt.Println(ans)
	if ans == "[]" {
		return nil
	}
	return &ans
}

func ResolveJsonAppData(data *string) (*map[string]string, error) {
	map1 := make(map[string]interface{})
	json.Unmarshal([]byte(*data), &map1)
	map2 := make(map[string]string, len(map1))
	for k, v := range map1 {
		map2[k] = fmt.Sprint(v)
	}
	if !JudgeAppData(&map2) {
		return nil, errors.New("Wrong App Data")
	}
	//fmt.Println(map2)
	return &map2, nil
}

func ResolveJsonRuleData(data *map[string]interface{}, check bool) (*map[string]string, *[]string, error) {
	// data := make(map[string]interface{})
	// json.Unmarshal([]byte(*data), &map1)
	listValue, ok := (*data)["device_id_list"].([]interface{})
	var keyStringValues []string
	if ok {
		keyStringValues = make([]string, len(listValue))
		for i, arg := range listValue {
			keyStringValues[i] = arg.(string)
		}
		delete(*data, "device_id_list")
	}
	map2 := make(map[string]string, len(*data))
	for k, v := range *data {
		map2[k] = fmt.Sprint(v)
	}
	//fmt.Println(keyStringValues)
	//fmt.Println(map2)
	if check && !JudgeLegalRule(&map2) {
		return nil, nil, errors.New("Wrong Rule Data")
	}
	return &map2, &keyStringValues, nil
}

func JudgeLegalRule(rule *map[string]string) bool {
	if ((*rule)["platform"] != "" &&
		strings.ToLower((*rule)["platform"]) != "ios" && strings.ToLower((*rule)["platform"]) != "android") ||
		((*rule)["cpu_arch"] != "32" && (*rule)["cpu_arch"] != "64" && (*rule)["cpu_arch"] != "") {
		LogMsg("platform not match")

		return false
	}

	vcode1 := (*rule)["update_version_code"]
	reg := regexp.MustCompile(`^\d+(\.\d+)*$`)
	if !reg.MatchString(vcode1) {
		LogfMsg("invalid update version code %v", vcode1)
		return false
	}
	vcode2 := (*rule)["max_update_version_code"]
	if !reg.MatchString(vcode2) {
		LogfMsg("invalid update version code %v", vcode2)
		return false
	}
	vcode3 := (*rule)["min_update_version_code"]
	if !reg.MatchString(vcode3) {
		LogfMsg("invalid update version code %v", vcode3)
		return false
	}
	for _, r := range (*rule)["min_os_api"] {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	for _, r := range (*rule)["max_os_api"] {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	// for i, r := range (*rule)["update_version_code"] {
	// 	if !unicode.IsDigit(r) && r != '.' {
	// 		LogMsg("illegal upvcode:", r)

	// 		return false
	// 	}
	// 	if i > 0 && (*rule)["update_version_code"][i-1] == '.' && r == '.' {
	// 		LogMsg("Consequence .")

	// 		return false
	// 	}
	// }
	// for i, r := range (*rule)["min_update_version_code"] {
	// 	if !unicode.IsDigit(r) && r != '.' {
	// 		LogMsg("illegal minupvcode:", r)

	// 		return false
	// 	}
	// 	if i > 0 && (*rule)["min_update_version_code"][i-1] == '.' && r == '.' {
	// 		LogMsg("Consequence .")

	// 		return false
	// 	}
	// }
	// for i, r := range (*rule)["max_update_version_code"] {
	// 	if !unicode.IsDigit(r) && r != '.' {
	// 		LogMsg("illegal max upvcode:", r)

	// 		return false
	// 	}
	// 	if i > 0 && (*rule)["max_update_version_code"][i-1] == '.' && r == '.' {
	// 		LogMsg("Consequence .")

	// 		return false
	// 	}
	// }
	// for i, r := range (*rule)["min_os_api"] {
	// 	if !unicode.IsDigit(r) && r != '.' {
	// 		LogMsg("illegal min osapi:", r)

	// 		return false
	// 	}
	// 	if i > 0 && (*rule)["min_os_api"][i-1] == '.' && r == '.' {
	// 		LogMsg("Consequence .")

	// 		return false
	// 	}
	// }
	// for i, r := range (*rule)["max_os_api"] {
	// 	if !unicode.IsDigit(r) && r != '.' {
	// 		LogMsg("illegal max osapi:", r)

	// 		return false
	// 	}
	// 	if i > 0 && (*rule)["max_os_api"][i-1] == '.' && r == '.' {
	// 		LogMsg("Consequence .")

	// 		return false
	// 	}
	// }
	//if ((*rule)["min_update_version_code"] != "" && (*rule)["max_update_version_code"] != "" &&
	//	VersionCmp((*rule)["min_update_version_code"], (*rule)["max_update_version_code"]) == 1) ||
	//	((*rule)["min_update_version_code"] != "" && (*rule)["max_update_version_code"] == "") ||
	//	((*rule)["min_update_version_code"] == "" && (*rule)["max_update_version_code"] != "") {
	//	LogMsg("更新版本範圍有誤")
	//
	//	return false
	//}
	if !((*rule)["min_update_version_code"] != "" && (*rule)["max_update_version_code"] != "" &&
		VersionCmp((*rule)["min_update_version_code"], (*rule)["max_update_version_code"]) <= 0) {
		LogMsg("更新版本範圍有誤")

		return false
	}
	//LogfMsg("%v %v %v", (*rule)["min_os_api"], (*rule)["max_os_api"], strings.Compare((*rule)["min_os_api"], (*rule)["max_os_api"]))
	//if ((*rule)["min_os_api"] != "" && (*rule)["max_os_api"] != "" &&
	//	strings.Compare((*rule)["min_os_api"], (*rule)["max_os_api"]) > 0) ||
	//	((*rule)["min_os_api"] != "" && (*rule)["max_os_api"] == "") ||
	//	((*rule)["min_os_api"] == "" && (*rule)["max_os_api"] != "") {
	//	LogMsg("不符合的api")
	//
	//	return false
	//}
	min_os,_ := strconv.Atoi((*rule)["min_os_api"])
	max_os,_ := strconv.Atoi((*rule)["max_os_api"])
	if !((*rule)["min_os_api"] != "" && (*rule)["max_os_api"] != "" &&
		min_os <= max_os ) {
		LogMsg("不符合的api")

		return false
	}
	return true
}
func JudgeAppData(rule *map[string]string) bool {
	if ((*rule)["device_platform"] != "" &&
		strings.ToLower((*rule)["device_platform"]) != "ios" && strings.ToLower((*rule)["device_platform"]) != "android") ||
		((*rule)["cpu_arch"] != "32" && (*rule)["cpu_arch"] != "64" && (*rule)["cpu_arch"] != "") {

		return false
	}
	_, ok := (*rule)["os_api"]
	if (strings.ToLower((*rule)["device_platform"]) == "ios" && ok) ||
		(strings.ToLower((*rule)["device_platform"]) == "android" && !ok) {

		return false
	}
	reg := regexp.MustCompile(`^\d+(\.\d+)*$`)
	vcode1 := (*rule)["update_version_code"]
	if !reg.MatchString(vcode1) {
		LogfMsg("invalid update version code %v", vcode1)
		return false
	}
	for _, r := range (*rule)["os_api"] {
		if !unicode.IsDigit(r) {

			return false
		}
	}
	for _, r := range (*rule)["aid"] {
		if !unicode.IsDigit(r) {

			return false
		}
	}
	return true
}
