package tools

import (
	"log"
	"strconv"
	"strings"
)

func versionCmp(s1 string, s2 string, flag bool) bool {
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
