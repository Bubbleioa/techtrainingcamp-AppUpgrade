// 对规则集的测试
package model

import (
	"fmt"
	"testing"
)

// 测试规则
func TestRules(t *testing.T) {
	r, e := resolveJson("rules.json")
	if e != nil {
		fmt.Println(e.Error())
	}
	for _, v := range r {
		m, e := RuleFactory(&v)
		fmt.Println(m)
		fmt.Println(e)
	}
}

// 测试获取所有规则
func TestAllRules(t *testing.T) {
	q := *GetAllRules()
	fmt.Println(q)
}
