// JSON解析器

package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 规则集待命中的一条规则
type Rule struct {
	UpdateVersionCode string      // 更新版本号
	MD5               string      // MD5
	DownloadURL       string      // 下载链接
	Title             string      // 更新标题
	UpdateTips        string      // 更新提示
	Conditions        []Condition // 条件集合
}

// 命中条件
type Condition struct {
	Type      string        // 匹配类型（精确匹配、范围匹配、数据库查找）
	Field     string        // 目标字段（和前端接口匹配）
	Content   string        // 精确匹配时匹配的内容
	ValueType string        // 内容的类型（int/string/valuetype），用来分派
	Rangev    []interface{} // 范围匹配时匹配的范围
}

// 使用Unmarshal解析，动态绑定到结构体字段
func resolveJson(fileName string) ([]Rule, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	bytevalue, _ := ioutil.ReadAll(jsonFile)
	rules := []Rule{}
	err = json.Unmarshal(bytevalue, &rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}
