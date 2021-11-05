package database

import (
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	InitClient()
	r := map[string]string{
		"aid":                     "200",
		"platform":                "iOS",
		"download_count":          "0",
		"hit_count":               "0",
		"download_url":            "http://baidu.com",
		"update_version_code":     "1.1.1",
		"md5":                     "123",
		"max_update_version_code": "1.1.0",
		"min_update_version_code": "1.0.0",
		"max_os_api":              "0",
		"min_os_api":              "0",
		"cpu_arch":                "32",
		"channel":                 "App Store",
		"title":                   "Update",
		"update_tips":             "yes",
		"enabled":                 "1",
		"create_date":             "123",
	}
	var wl []string
	wl = append(wl, "1")
	wl = append(wl, "2")
	wl = append(wl, "3")
	tmp_id := cur_id
	AddRule(r, wl)
	//fmt.Println(err)

	val, _ := CheckAppidInWhiteList(strconv.Itoa(tmp_id)+"set", "1")
	if val == false {
		t.Errorf("UnExpected!")
	}

}

func TestGetRuleAtt(t *testing.T) {
	InitClient()
	val, _ := GetRuleAttr(strconv.Itoa(cur_id-1), "aid")
	if val != "200" {
		t.Errorf("UnExpected!%v %v", cur_id, val)
	}
}
