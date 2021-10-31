package model

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
