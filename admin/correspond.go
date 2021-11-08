package admin

import (
	"net/http"
	"techtrainingcamp-AppUpgrade/database"
	"techtrainingcamp-AppUpgrade/tools"

	"github.com/gin-gonic/gin"
)

func QueryAllRules(c *gin.Context) {
	// lst, e := database.QueryAllRules()
	lst, e := query_allrules_testbench()
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": e.Error()})
		return
	}
	str := tools.ConvertSimplifiedRulesListToJson(lst)
	c.String(http.StatusOK, *str)
}

func QueryRule(c *gin.Context) {
	ruleid := c.Query("ruleid")
	// ml, lst, e := database.QueryRuleByID(ruleid)
	ml, lst, e := queryrulebyid_testbench(ruleid)
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": e.Error()})
	}
	if len(*ml) > 1 {
		c.JSON(http.StatusBadGateway, gin.H{"message": "illegle ruleid"})
	}
	if len(*ml) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"messgae": "no such rule"})
	}
	str := tools.ConvertFullRuleToJSON(&((*ml)[0]), lst)
	c.String(http.StatusOK, *str)
}

func UpdateRule(c *gin.Context) {
	mp := make(map[string]interface{})
	c.BindJSON(&mp)
	mm, lst, e := tools.ResolveJsonRuleData(&mp, false)
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": "Illegal rule data"})
	}
	var v string
	var ok bool
	v, ok = (*mm)["ruleid"]
	if !ok {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": "No ruleid!"})
		return
	}
	// oldrules, oldlst, e := database.QueryRuleByID(v)
	oldrules, oldlst, e := queryrulebyid_testbench(v)
	oldrule := (*oldrules)[0]
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": "an error occurred"})
		return
	}
	for idx, ele := range *mm {
		_, ok := oldrule[idx]
		if !ok {
			c.JSON(http.StatusBadGateway, gin.H{"messgae": "Invalid value : (" + idx + "," + ele + ") "})
			return
		}
		oldrule[idx] = ele
	}
	if !tools.JudgeLegalRule(&oldrule) {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": "Invalid value"})
		return
	}
	if lst != nil {
		// e = database.UpdateRule(&oldrule, lst)
		e = update_database_testbench(&oldrule, lst)
	} else {
		// e = database.UpdateRule(&oldrule, oldlst)
		e = update_database_testbench(&oldrule, oldlst)
	}
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": "Data insert error..."})
		return
	}
	c.Status(http.StatusOK)
}

func CreateRule(c *gin.Context) {
	mp := make(map[string]interface{})
	c.BindJSON(&mp)
	mm, lst, e := tools.ResolveJsonRuleData(&mp, true)
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": "Illegal rule data"})
	}
	e = database.AddRule(mm, lst)
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": "Data insert error..."})
	}
	c.Status(http.StatusOK)
}

func DeleteRule(c *gin.Context) {
	ruleid := c.Query("ruleid")
	e := database.DeleteRule(ruleid)
	if e != nil {
		c.JSON(http.StatusBadGateway, gin.H{"messgae": e.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}
