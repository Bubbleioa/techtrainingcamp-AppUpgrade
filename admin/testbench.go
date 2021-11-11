package admin

func query_allrules_testbench() (*[]map[string]string, error) {
	mp := make([]map[string]string, 3)
	for i, _ := range mp {
		mp[i] = map[string]string{
			"ruleid":              "10",
			"aid":                 "3",
			"platform":            "android",
			"update_version_code": "6.2.1",
		}
	}
	return &mp, nil
}

func queryrulebyid_testbench(ruleid string) (*[]map[string]string, *[]string, error) {
	mp := map[string]string{
		"ruleid":                  "10",
		"aid":                     "3",
		"platform":                "android",
		"update_version_code":     "6.2.1",
		"download_url":            "baidu.com",
		"md5":                     "dsfsdacad",
		"max_update_version_code": "5.8.2",
		"min_update_version_code": "5.1.1",
		"max_os_api":              "21",
		"min_os_api":              "18",
		"cpu_arch":                "32",
		"channel":                 "华为应用市场",
		"title":                   "升级啦",
		"update_tips":             "快升级",
	}
	lst := []string{
		"asdasasc", "adfcc", "dfsv",
	}
	ml := []map[string]string{mp}
	return &ml, &lst, nil
}

func update_database_testbench(rulemap *map[string]string, devicelst *[]string) error {
	return nil
}
