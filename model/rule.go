package model

import "fmt"

// 返回类型的映射
type Response map[string]string

// 规则集的映射
type Conditionset map[string]IMatcher

// 单条规则
type SingleRule struct {
	Res Response
	Hit Conditionset
}

// 获取所有规则
func GetAllRules() *[]SingleRule {
	rules, err := resolveJson("config/rules.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	ruleset := []SingleRule{}
	for _, v := range rules {
		rule := SingleRule{}
		rule.Res = make(Response)
		rule.Hit = make(Conditionset)
		rule.Res["update_version_code"] = v.UpdateVersionCode
		rule.Res["md5"] = v.MD5
		rule.Res["download_url"] = v.DownloadURL
		rule.Res["title"] = v.Title
		rule.Res["update_tips"] = v.UpdateTips
		rule.Hit, err = RuleFactory(&v)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		ruleset = append(ruleset, rule)
	}
	return &ruleset
}
