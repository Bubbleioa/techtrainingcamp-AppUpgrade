package model

import "fmt"

type Rule struct {
	// rule confition
	MinUpdateVersionCode string `json:"min_update_version_code"`
	MaxUpdateVersionCode string `json:"max_update_version_code"`
	MinOsApi             int    `json:"min_os_api"`
	MaxOsApi             int    `json:"max_os_api"`
	Platform             string `json:"platform"`
	CpuArch              string `json:"cpu_arch"`
	Channel              string `json:"channel"`
	// apk or ipk link
	DownloadUrl       string `json:"download_url"`
	UpdateVersionCode string `json:"update_version_code"`
	Md5               string `json:"md5"`
	Title             string `json:"title"`
	UpdateTips        string `json:"update_tips"`
}
func GetAllRules() *[]Rule {
	rules := []Rule{}

	rules = append(rules, Rule{
		MinUpdateVersionCode: "8.4.0",
		MaxUpdateVersionCode: "8.8.8",
		MinOsApi:             10,
		MaxOsApi:             20,
		Platform:             "Android",
		CpuArch:              "32",
		Channel:              "dsd",
		DownloadUrl:          "https://baidu1.com",
		UpdateVersionCode:    "4.1",
		Md5:                  "dsafaf",
		Title:                "dsfds",
		UpdateTips:           "fdsafdas",
	})

	rules = append(rules, Rule{
		MinUpdateVersionCode: "8.4.0",
		MaxUpdateVersionCode: "8.8.8",
		MinOsApi:             10,
		MaxOsApi:             20,
		Platform:             "Android",
		CpuArch:              "64",
		Channel:              "dsd",
		DownloadUrl:          "https://baidu2.com",
		UpdateVersionCode:    "4.1",
		Md5:                  "dsafaf",
		Title:                "dsfds",
		UpdateTips:           "fdsafdas",
	})

	rules = append(rules, Rule{
		MinUpdateVersionCode: "8.4.0",
		MaxUpdateVersionCode: "8.8.8",
		MinOsApi:             10,
		MaxOsApi:             20,
		Platform:             "iOS",
		CpuArch:              "32",
		Channel:              "dsd",
		DownloadUrl:          "https://baidu3.com",
		UpdateVersionCode:    "4.1",
		Md5:                  "dsafaf",
		Title:                "dsfds",
		UpdateTips:           "fdsafdas",
	})

	rules = append(rules, Rule{
		MinUpdateVersionCode: "8.4.0",
		MaxUpdateVersionCode: "8.8.8",
		MinOsApi:             10,
		MaxOsApi:             20,
		Platform:             "iOS",
		CpuArch:              "64",
		Channel:              "dsd",
		DownloadUrl:          "https://baidu4.com",
		UpdateVersionCode:    "4.1",
		Md5:                  "dsafaf",
		Title:                "dsfds",
		UpdateTips:           "fdsafdas",
	})

	return &rules
}

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
func GetRulesFromFile() *[]SingleRule {
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
