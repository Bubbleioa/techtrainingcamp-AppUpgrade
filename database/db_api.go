package database

import (
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var BytesKind = reflect.TypeOf(sql.RawBytes{}).Kind()
var TimeKind = reflect.TypeOf(mysql.NullTime{}).Kind()

func checkErr(err error) {
	if err != nil {
		fmt.Printf("checkErr:%v", err)
	}
}

func ToStr(strObj interface{}) string {
	switch v := strObj.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", strObj)
	}
}

func ToInt(intObj interface{}) int {
	// 假定int == int64，运行在64位机
	switch v := intObj.(type) {
	case []byte:
		return ToInt(string(v))
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		if v > math.MaxInt64 {
			info := fmt.Sprintf("ToInt, error, overflowd %v", v)
			panic(info)
		}
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		strv := v
		if strings.Contains(v, ".") {
			strv = strings.Split(v, ".")[0]
		}
		if strv == "" {
			return 0
		}
		if intv, err := strconv.Atoi(strv); err == nil {
			return intv
		}
	}
	fmt.Printf(fmt.Sprintf("ToInt err, %v, %v not supportted\n", intObj,
		reflect.TypeOf(intObj).Kind()))
	return 0
}

func CheckDeviceIDInWhiteList(ruleid string, userid string) (bool, error) {
	c, err := RedisCheckWhiteList(ruleid, userid)
	checkErr(err)
	return c, err
}

func GetRuleAtt(ruleid string, field string) (string, error) {
	return RedisGetRuleAttr(ruleid, field)
}

// func UpdateUserDownloadStatus(ruleid string, status bool) error {
// 	err:=RedisUpdateDownloadStatus(ruleid,status)
// 	checkErr(err)

// }

//查询所有规则，为了保证完整性，对 mysql 查询
func QueryAllRules() (*[]map[string]string, error) {
	val, _, err := MysqlQueryRules("0")
	return val, err
}

//优先对 redis 查询，若没查询到，对 mysql 查询并更新 redis
func QueryRuleByID(ruleid string) (*[]map[string]string, *[]string, error) {
	res, devices, err := RedisQueryRuleByID(ruleid)
	if err != nil {
		fmt.Println("Redis not found, query mysql next...")
	} else {
		return res, devices, err
	}
	res, devices, err = MysqlQueryRules(ruleid)
	if err != nil {
		fmt.Println("Wrong ID!")
	}
	RedisUpdateRuleWithList(ruleid, (*res)[0])
	return res, devices, err
}

//提供一个 string-string 的哈希表和白名单，向 mysql 添加规则。
func AddRule(rulemap *map[string]string, devicelst *[]string) error {
	//fmt.Println((*rulemap)["id"])
	err := MysqlAddRule(rulemap, devicelst)
	checkErr(err)
	err = RedisUpdateRule((*rulemap)["id"], *rulemap, *devicelst)
	checkErr(err)
	return err
}

func UpdateRule(rulemap *map[string]string, devicelst *[]string) error {
	err := RedisUpdateRule((*rulemap)["id"], *rulemap, *devicelst)
	checkErr(err)
	err = MysqlUpdateRule(rulemap, devicelst)
	checkErr(err)
	return err
}

func DeleteRule(ruleid string) error {
	err := MysqlDeleteRule(ruleid)
	checkErr(err)
	err = RedisDeleteRule(ruleid)
	checkErr(err)
	return err
}

// 这个接口直接放在了 mysql.go 中
// func GetDownloadRatio(ruleid string) (float64, error)
