package tools

import (

	"testing"
)

func TestVersionCmp(t *testing.T) {

	m := make(map[string]string)
	m["update_version_code"] = "8.8.8"
	m["max_update_version_code"] = "8.8.8"
	m["min_update_version_code"] = "8.8.8"
	m["min_os_api"] = "8"
	m["max_os_api"] = "8"
	if JudgeLegalRule(&m) == false {
		t.Errorf("error")
	}
}
