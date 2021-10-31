// 判断匹配的接口

package model

import (
	"errors"
	"strings"

	"github.com/spf13/cast"
)

// IComparator接口，用来实现不同类型的比较
type IComparator interface {
	// 相等
	Eq(other IComparator) (bool, error)
	// 小于等于
	Le(other IComparator) (bool, error)
	// 大于等于
	Ge(other IComparator) (bool, error)
	// 实现动态绑定类型的trick，不知道有没有更好的解决办法
	CastStr(b string) IComparator
}

// 实现了特殊比较方法的int
type MyInt int

func (a MyInt) Eq(b IComparator) (bool, error) {
	s, ok := b.(MyInt)
	if !ok {
		return false, errors.New("type error")
	} else {
		return int(a) == int(s), nil
	}
}

func (a MyInt) Le(b IComparator) (bool, error) {
	s, ok := b.(MyInt)
	if !ok {
		return false, errors.New("type error")
	} else {
		return int(a) <= int(s), nil
	}
}

func (a MyInt) Ge(b IComparator) (bool, error) {
	s, ok := b.(MyInt)
	if !ok {
		return false, errors.New("type error")
	} else {
		return int(a) >= int(s), nil
	}
}

func (a MyInt) CastStr(b string) IComparator {
	return MyInt(cast.ToInt(b))
}

// 实现了特殊比较方法的string
type MyString string

func (a MyString) Eq(b IComparator) (bool, error) {
	s, ok := b.(MyString)
	if !ok {
		return false, errors.New("type error")
	} else {
		return string(a) == string(s), nil
	}
}

func (a MyString) Le(b IComparator) (bool, error) {
	s, ok := b.(MyString)
	if !ok {
		return false, errors.New("type error")
	} else {
		return string(a) <= string(s), nil
	}
}

func (a MyString) Ge(b IComparator) (bool, error) {
	s, ok := b.(MyString)
	if !ok {
		return false, errors.New("type error")
	} else {
		return string(a) >= string(s), nil
	}
}

func (a MyString) CastStr(b string) IComparator {
	return MyString(b)
}

// 实现了特殊比较方法的版本号类型
type VersionID struct {
	VersionStr string
}

func (a VersionID) Eq(b IComparator) (bool, error) {
	s, ok := b.(VersionID)
	if !ok {
		return false, errors.New("type error")
	} else {
		arr1 := strings.Split(a.VersionStr, ".")
		arr2 := strings.Split(s.VersionStr, ".")
		for i := len(arr1) - 1; i >= 0; i = i - 1 {
			if cast.ToInt(arr1[i]) == 0 {
				arr1 = arr1[0:i]
			} else {
				break
			}
		}
		for i := len(arr2) - 1; i >= 0; i = i - 1 {
			if cast.ToInt(arr2[i]) == 0 {
				arr2 = arr2[0:i]
			} else {
				break
			}
		}
		if len(arr1) != len(arr2) {
			return false, nil
		}
		for id := range arr1 {
			if cast.ToInt(arr1[id]) != cast.ToInt(arr2[id]) {
				return false, nil
			}
		}
		return true, nil
	}
}

// 局部取int的min函数
func intMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (a VersionID) Le(b IComparator) (bool, error) {
	s, ok := b.(VersionID)
	if !ok {
		return false, errors.New("type error")
	} else {
		arr1 := strings.Split(a.VersionStr, ".")
		arr2 := strings.Split(s.VersionStr, ".")
		for i := len(arr1) - 1; i >= 0; i = i - 1 {
			if cast.ToInt(arr1[i]) == 0 {
				arr1 = arr1[0:i]
			} else {
				break
			}
		}
		for i := len(arr2) - 1; i >= 0; i = i - 1 {
			if cast.ToInt(arr2[i]) == 0 {
				arr2 = arr2[0:i]
			} else {
				break
			}
		}
		for i := 0; i < intMin(len(arr1), len(arr2)); i = i + 1 {
			if cast.ToInt(arr1[i]) < cast.ToInt(arr2[i]) {
				return true, nil
			}
			if cast.ToInt(arr1[i]) > cast.ToInt(arr2[i]) {
				return false, nil
			}
		}
		if len(arr1) <= len(arr2) {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func (a VersionID) Ge(b IComparator) (bool, error) {
	s, ok := b.(VersionID)
	if !ok {
		return false, errors.New("type error")
	} else {
		arr1 := strings.Split(a.VersionStr, ".")
		arr2 := strings.Split(s.VersionStr, ".")
		for i := len(arr1) - 1; i >= 0; i = i - 1 {
			if cast.ToInt(arr1[i]) == 0 {
				arr1 = arr1[0:i]
			} else {
				break
			}
		}
		for i := len(arr2) - 1; i >= 0; i = i - 1 {
			if cast.ToInt(arr2[i]) == 0 {
				arr2 = arr2[0:i]
			} else {
				break
			}
		}
		for i := 0; i < intMin(len(arr1), len(arr2)); i = i + 1 {
			if cast.ToInt(arr1[i]) < cast.ToInt(arr2[i]) {
				return false, nil
			}
			if cast.ToInt(arr1[i]) > cast.ToInt(arr2[i]) {
				return true, nil
			}
		}
		if len(arr1) >= len(arr2) {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func (a VersionID) CastStr(b string) IComparator {
	return VersionID{b}
}

// IMatcher接口，用来判断某个对象是否符合条件
type IMatcher interface {
	// 相同类型的成功
	Success(in IComparator) (bool, error)
	// 由于传入都是string，所以需要特殊判断
	SuccessStr(in string) (bool, error)
}

// 精确匹配
type Equivalence struct {
	TargetVal IComparator
}

func (eq Equivalence) Success(in IComparator) (bool, error) {
	return in.Eq(eq.TargetVal)
}

func (eq Equivalence) SuccessStr(in string) (bool, error) {
	ic := eq.TargetVal.CastStr(in)
	return eq.Success(ic)
}

// 范围匹配
type InRange struct {
	MinVal IComparator
	MaxVal IComparator
}

func (ir InRange) Success(in IComparator) (bool, error) {
	ge, e1 := in.Ge(ir.MinVal)
	le, e2 := in.Le(ir.MaxVal)
	if e1 != nil {
		return false, e1
	}
	if e2 != nil {
		return false, e2
	}
	return ge && le, nil
}

func (ir InRange) SuccessStr(in string) (bool, error) {
	ic := ir.MinVal.CastStr(in)
	return ir.Success(ic)
}

// TODO：数据库
