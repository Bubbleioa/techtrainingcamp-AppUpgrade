// 已经废弃的解决方案
package model

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// )

// // type Rules struct {
// // 	XMLName xml.Name     `xml:"rules"`
// // 	Rules   []AppVersion `xml:"appversion"`
// // }

// type AppVersion struct {
// 	// XMLName           xml.Name    `xml:"appversion"`
// 	UpdateVersionCode string      //`xml:"update-version-code"`
// 	Md5               string      //`xml:"md5"`
// 	DownloadURL       string      //`xml:"download-url"`
// 	Title             string      //`xml:"title"`
// 	UpdateTip         string      //`xml:"update-tips"`
// 	Conditions        []Condition //`xml:"condition"`
// }

// type Condition struct {
// 	//XMLName   xml.Name //`xml:"condition"`
// 	Type      string //`xml:"type,attr"`
// 	Field     string //`xml:"field,attr"`
// 	ValueType string //`xml:"valuetype,attr"`
// 	MinValue  string //`xml:"min-value"`
// 	MaxValue  string //`xml:"max-value"`
// }

// func printElmt(s string, depth int) {
// 	for n := 0; n < depth; n++ {
// 		fmt.Print("  ")
// 	}
// 	fmt.Println(s, " ")
// }

// func loadXMLFile(fileName string) {
// 	xmlFile, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer xmlFile.Close()

// 	bytevalue, _ := ioutil.ReadAll(xmlFile)
// 	r := strings.NewReader(string(bytevalue))
// 	parser := xml.NewDecoder(r)
// 	dep := 0
// 	var versions []AppVersion
// 	var version AppVersion
// 	var conditions []Condition
// 	var condition Condition
// 	var currTagName string
// 	for {
// 		token, err := parser.Token()
// 		if err != nil {
// 			break
// 		}
// 		switch t := token.(type) {
// 		case xml.StartElement:
// 			// printElmt(xml.StartElement(t).Name.Local, dep)
// 			// fmt.Println(xml.StartElement(t).Attr)
// 			dep++
// 			currTagName = xml.StartElement(t).Name.Local
// 			switch dep {
// 			case 1:
// 				version = AppVersion{}
// 			case 2:
// 				if currTagName == "conditions" {
// 					conditions = []Condition{}
// 				}
// 			case 3:
// 				attr := xml.StartElement(t).Attr
// 				for _, v := range attr {
// 					if v.Name.Local == "type" {
// 						condition.
// 					}
// 				}
// 			}
// 		case xml.EndElement:
// 			dep--
// 			// printElmt(xml.EndElement(t).Name.Local, dep)
// 			switch dep {
// 			case 1:
// 				versions = append(versions, version)
// 			}
// 		case xml.CharData:
// 			bytes := xml.CharData(t)
// 			str := string([]byte(bytes))
// 			if dep == 2 {
// 				switch currTagName {
// 				case "update-version-code":
// 					version.UpdateVersionCode = str
// 				case "md5":
// 					version.Md5 = str
// 				case "download-url":
// 					version.DownloadURL = str
// 				case "title":
// 					version.Title = str
// 				case "update-tips":
// 					version.UpdateTip = str
// 				}
// 			}
// 		}
// 	}
// 	// fmt.Println(rules.XMLName)
// }
