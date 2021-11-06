package database

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenMysql() error {
	var err error
	db, err = sql.Open("mysql", "test:123456@/app")
	if err != nil {
		fmt.Println("数据库链接错误", err)
	}
	//延迟到函数结束关闭链接
	//defer db.Close()
	return err
}

func MysqlAddRule(rulemap *map[string]string, devicelst *[]string) error {
	OpenMysql()
	defer db.Close()
	var devices string
	for _, val := range *devicelst {
		devices += val + ","
	}
	_, err := db.Exec("insert into rules(aid,platform,download_url,update_version_code,device_list,md5,max_update_version_code,min_update_version_code,max_os_api,min_os_api,cpu_arch,channel,title,update_tips) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)", (*rulemap)["aid"], (*rulemap)["platform"], (*rulemap)["download_url"], (*rulemap)["update_version_code"], devices, (*rulemap)["md5"], (*rulemap)["max_update_version_code"], (*rulemap)["min_update_version_code"], (*rulemap)["max_os_api"], (*rulemap)["min_os_api"], (*rulemap)["cpu_arch"], (*rulemap)["channel"], (*rulemap)["title"], (*rulemap)["update_tips"])
	if err != nil {
		panic(err)
	}

	return err
}

func RowsToMap(rows *sql.Rows) *[]map[string]interface{} {
	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		cols, err := rows.Columns()
		checkErr(err)

		colsTypes, err := rows.ColumnTypes()
		checkErr(err)

		dest := make([]interface{}, len(cols))
		destPointer := make([]interface{}, len(cols))
		for i := range dest {
			destPointer[i] = &dest[i]
		}

		err = rows.Scan(destPointer...)
		checkErr(err)

		rowResult := make(map[string]interface{})
		for i, colVal := range dest {
			colName := cols[i]
			itemType := colsTypes[i].ScanType()
			//fmt.Printf("type %v \n", itemType)

			switch itemType.Kind() {
			case BytesKind:
				rowResult[colName] = ToStr(colVal)

			case reflect.Int, reflect.Int8,
				reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

				rowResult[colName] = ToInt(colVal)

			case TimeKind:
				fmt.Println("time", colVal, reflect.TypeOf(colVal))
				rowResult[colName] = ToStr(colVal)
			default:
				rowResult[colName] = ToStr(colVal)
			}
		}
		result = append(result, rowResult)
	}
	return &result
}

//根据id查询规则，"0"代表全部
func MysqlQueryRules(ruleid string) (*[]map[string]interface{}, error) {
	OpenMysql()
	defer db.Close()
	if ruleid == "0" {
		dbrows, err := db.Query("select * from rules")
		if err != nil {
			panic(err)
			// return nil, err
		}
		return RowsToMap(dbrows), err
	} else {
		dbrows, err := db.Query("select * from rules where id=?", ruleid)
		if err != nil {
			panic(err)
			// return nil, err
		}
		return RowsToMap(dbrows), err
	}
}
