package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// data Source: https://docs.google.com/spreadsheets/d/1I9EXxe-pWLhcLosakg5TPt98ERY6tdpJn1KngIGY7oY/edit#gid=0
// script:
// curl 'https://raw.githubusercontent.com/kelvin2go/covid19-tw/master/latest.json' > downloads/patients_raw.json
// go run tw_g0v_google_sheet_json_to_pat_data.go > downloads/tmp.json
// url: https://raw.githubusercontent.com/kelvin2go/covid19-tw/master/latest.json
// url: https://sheets.googleapis.com/v4/spreadsheets/1I9EXxe-pWLhcLosakg5TPt98ERY6tdpJn1KngIGY7oY/values/%E7%A2%BA%E8%A8%BA%E7%97%85%E4%BE%8B!A1:L?majorDimension=ROWS&key={key}

// https://siongui.github.io/2015/02/27/go-parse-rss2/

type SourceFormat struct {
	Date        string
	ClassAToday string // 疑新送驗
	ClassATotal string // 疑新送驗累計
	ClassBToday string // 居檢送驗
	ClassBTotal string // 居檢送驗累計
	ClassCToday string // 武肺通報
	ClassCTotal string // 武肺通報累計
}

// "日期",
//      "疑新送驗",
//      "疑新送驗累計",
//      "居檢送驗",
//      "居檢送驗累計",
//      "武肺通報",
//      "武肺通報累計",
//      "檢驗人數",
//      "檢驗人數累計",
//      "確診人數",
//      "確診人數累計",
//      "單日陽性率"

// type DestinateFormat struct {
// 	Id            string `json:"id"`
// 	ReleaseDate   string `json:"リリース日"`
// 	MoveInDate    string `json:"移入日"`
// 	Place         string `json:"居住地"`
// 	RelationPlace string `json:"相關地點"`
// 	Imported      string `json:"境外或本土"`
// 	Age           string `json:"年代"`
// 	Gender        string `json:"性別"`
// 	IsExit        string `json:"退院"`
// 	Link          string `json:"link"`
// }

// {
//     "id": "1",
//     "リリース日": "2020-01-21T08:00:00.000Z",
//     "移入日": "2020-01-20T08:00:00.000Z",
//     "居住地": "高雄市",
//     "年代": "55代",
//     "性別": "女性",
//     "退院": "〇"
// },

// "CDC 新聞稿 URL",
// "案例編號",
// "性別",
// "年齡層\n(~多歲)",
// "就醫日",
// "確診日",
// "其他公布說明",
// "死亡日",
// "解除隔離日"

func main() {
	// fmt.Println("start")
	file, err := os.Open("./downloads/inspections_summary_raw.json")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	fileData, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := map[string]interface{}{}
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	values := data["values"].([]interface{})
	title := values[0].([]interface{})
	rows := values[1:]
	// fmt.Println("title", title)

	keyToRowNumber := map[string]int{}
	for v, key := range title {
		keyToRowNumber[key.(string)] = v
	}

	srcList := []SourceFormat{}

	// CaseId           string
	// Gender           string
	// Age              string
	// DateForHospital  string
	// DateForConfirm   string
	// Note             string
	// DateForDeath     string
	// DateForDischarge string

	for _, item := range rows {
		it2 := item.([]interface{})
		newItem := SourceFormat{
			Date:        it2[keyToRowNumber["日期"]].(string),
			ClassAToday: it2[keyToRowNumber["疑新送驗"]].(string),
			ClassATotal: it2[keyToRowNumber["疑新送驗累計"]].(string),
			ClassBToday: it2[keyToRowNumber["居檢送驗"]].(string),
			ClassBTotal: it2[keyToRowNumber["居檢送驗累計"]].(string),
			ClassCToday: it2[keyToRowNumber["武肺通報"]].(string),
			ClassCTotal: it2[keyToRowNumber["武肺通報累計"]].(string),
		}

		srcList = append(srcList, newItem)
	}
	// fmt.Println(srcList)

	嚴重特殊傳染性肺炎通報List := []int{}
	居家檢疫送驗List := []int{}
	疑似新冠病毒感染送驗List := []int{}
	Label := []string{}

	for _, item := range srcList {
		v, _ := strconv.Atoi(strings.Replace(item.ClassAToday, ",", "", -1))
		疑似新冠病毒感染送驗List = append(疑似新冠病毒感染送驗List, v)
		v, _ = strconv.Atoi(strings.Replace(item.ClassBToday, ",", "", -1))
		居家檢疫送驗List = append(居家檢疫送驗List, v)
		v, _ = strconv.Atoi(strings.Replace(item.ClassCToday, ",", "", -1))
		嚴重特殊傳染性肺炎通報List = append(嚴重特殊傳染性肺炎通報List, v)
		Label = append(Label, item.Date)
	}

	// fmt.Println(data)
	// list := []map[string]interface{}{}
	// for _, item := range data.ItemList {
	// 	list = append(list, map[string]interface{}{
	// 		"date": item.Updated,
	// 		"text": item.Title,
	// 		"url":  item.Link,
	// 	})
	// }

	out := map[string]interface{}{}
	out["date"] = time.Now().Format("2006/01/02 15:04")
	out["data"] = map[string]interface{}{
		"疑似新冠病毒感染送驗":  疑似新冠病毒感染送驗List,
		"居家檢疫送驗":      居家檢疫送驗List,
		"嚴重特殊傳染性肺炎通報": 嚴重特殊傳染性肺炎通報List,
	}
	out["labels"] = Label

	outBytes, _ := json.MarshalIndent(out, "", "    ")
	fmt.Println(string(outBytes))

}
