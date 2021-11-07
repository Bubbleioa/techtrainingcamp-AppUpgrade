package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func VersionCmp(a string, b string) int {
	arr1 := strings.Split(a, ".")
	arr2 := strings.Split(b, ".")
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
		for index := 0; index < arrLen; index++ {
			if strings.Compare(arr1[index], arr2[index]) < 0 {
				return -1
			}else if strings.Compare(arr1[index], arr2[index]) > 0 {
				return 1
			}
		}
		return 0
}

func ConvertFullRuleToJSON(rule *map[string]string, devicelist *[]string) *string{
	mjson,_ :=json.Marshal(*rule)
	mString :=string(mjson)
	var mString2 string
	mString2 += "["
	for _, i := range *devicelist {
		mString2 += i + ","
	}
	mString2 = mString2[0 : len(mString2) - 1]
	mString2 += "]"
	mString += "," + mString2
	fmt.Println(mString)
	return &mString
}

func ConvertSimplifiedRulesListToJson(rules *[]map[string]string) *string{
	var ans string
	ans += "["
	for _, i := range *rules{
		mjson,_ :=json.Marshal(i)
		mString :=string(mjson)
		ans += mString + ","
	}
	ans = ans[0 : len(ans) - 1]
	ans += "]"
	fmt.Println(ans)
	return &ans
}