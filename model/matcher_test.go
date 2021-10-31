// 对Matcher的单元测试，覆盖率不太够

package model

import "testing"

func TestMyInt(t *testing.T) {
	var i IMatcher = Equivalence{MyInt(5)}
	if a, _ := i.Success(MyInt(5)); a == false {
		t.Errorf("UnExpected!")
	}
}

func TestMyString(t *testing.T) {
	var i IMatcher = InRange{
		MinVal: MyString("aaa"),
		MaxVal: MyString("bbb"),
	}
	if a, _ := i.Success(MyString("abb")); a == false {
		t.Errorf("UnExpected!")
	}
}

func TestVersion1(t *testing.T) {
	var v1 = VersionID{"8.4.2"}
	var v2 = VersionID{"8.1.3"}
	if a, _ := v1.Ge(v2); a == false {
		t.Errorf("UnExpected!")
	}
}

func TestVersion2(t *testing.T) {
	var v1 = VersionID{"8.4.2.1"}
	var v2 = VersionID{"8.4.2"}
	if a, _ := v1.Ge(v2); a == false {
		t.Errorf("UnExpected!")
	}
}

func TestVersion3(t *testing.T) {
	var v1 = VersionID{"8.4.2"}
	var v2 = VersionID{"8.04.2"}
	if a, _ := v1.Eq(v2); a == false {
		t.Errorf("UnExpected!")
	}
}

func TestVersion4(t *testing.T) {
	var v1 = VersionID{"8.4.2"}
	var v2 = VersionID{"8.4.2.0"}
	if a, _ := v1.Eq(v2); a == false {
		t.Errorf("UnExpected!")
	}
}

func TestVersion5(t *testing.T) {
	var v1 = VersionID{"8.4"}
	var v2 = VersionID{"8.1.3.4"}
	if a, _ := v1.Ge(v2); a == false {
		t.Errorf("UnExpected!")
	}
}
