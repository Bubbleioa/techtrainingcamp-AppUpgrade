package tools

import "testing"

func TestVersionCmp(t *testing.T) {
	var s1 = "8.0.0"
	var s2 = "8.0.0"
	if VersionCmp(s1, s2) != 0 {
		t.Errorf("error")
	}
}
