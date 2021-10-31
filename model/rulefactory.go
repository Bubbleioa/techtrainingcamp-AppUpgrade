// 用来产生规则的工厂函数
package model

import (
	"errors"
	"fmt"

	"github.com/spf13/cast"
)

func parseInt(x float64) int {
	return int(x + 1e-10)
}

// 根据输入类型和匹配类型动态产生IMatcher对象
func ConditionFactory(c *Condition) (IMatcher, error) {
	vtype := c.ValueType
	cmptype := c.Type
	if vtype == "" || cmptype == "" {
		return nil, errors.New("not valid type")
	}
	if cmptype == "equal" {
		content := c.Content
		switch vtype {
		case "int":
			nc := cast.ToInt(content)
			return Equivalence{MyInt(nc)}, nil
		case "string":
			return Equivalence{MyString(content)}, nil
		case "appversion":
			return Equivalence{VersionID{content}}, nil
		default:
			return nil, fmt.Errorf("not valid type: %v", vtype)
		}
	} else if cmptype == "range" {
		rng := c.Rangev
		if len(rng) != 2 {
			return nil, errors.New("valid range need 2 values")
		}
		switch vtype {
		case "int":
			l, ok1 := rng[0].(float64)
			r, ok2 := rng[1].(float64)
			if !ok1 || !ok2 {
				return nil, errors.New("cast failed")
			}
			return InRange{MyInt(parseInt(l)), MyInt(parseInt(r))}, nil
		case "string":
			l, ok1 := rng[0].(string)
			r, ok2 := rng[1].(string)
			if !ok1 || !ok2 {
				return nil, errors.New("cast failed")
			}
			return InRange{MyString(l), MyString(r)}, nil
		case "appversion":
			l, ok1 := rng[0].(string)
			r, ok2 := rng[1].(string)
			if !ok1 || !ok2 {
				return nil, errors.New("cast failed")
			}
			return InRange{VersionID{l}, VersionID{r}}, nil
		default:
			return nil, fmt.Errorf("not a valid type: %v", vtype)
		}
	} else {
		return nil, fmt.Errorf("not a valid type: %v", cmptype)
	}
}

// 把条件集合起来，生成从目标字段向IMatcher的映射
func RuleFactory(r *Rule) (map[string]IMatcher, error) {
	m := map[string]IMatcher{}
	for _, v := range r.Conditions {
		matcher, err := ConditionFactory(&v)
		if err != nil {
			return nil, err
		}
		field := v.Field
		if field == "" {
			return nil, errors.New("invalid field name")
		}
		m[field] = matcher
	}
	return m, nil
}
