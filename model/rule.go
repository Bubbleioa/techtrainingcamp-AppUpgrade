package model

type Rule struct {
	// rule confition
	MinVersion int `json:"min_version"`
	MaxVersion int `json:"max_version"`
	MinUserDID int `json:"min_user_did"`
	MaxUserDID int `json:"max_user_did"`

	// apk or ipk link
	GrayLink string `json:"gray_link"`
}

func GetAllRules() *[]Rule {
	rules := []Rule{}

	rules = append(rules, Rule{
		MinVersion: 10,
		MaxVersion: 20,
		MinUserDID: 10,
		MaxUserDID: 20,
		GrayLink:   "https://baidu.com",
	})
	rules = append(rules, Rule{
		MinVersion: 0,
		MaxVersion: 9,
		MinUserDID: 0,
		MaxUserDID: 9,
		GrayLink:   "https://baidu.com",
	})

	return &rules
}
