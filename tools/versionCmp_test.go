package tools

import "testing"

func TestVersionCmp(t *testing.T) {
	var s1 = "8.6.0"
	var s2 = "8.13.0"
	if VersionCmp(s1, s2) != -1 {
		t.Errorf("error")
	}
}
