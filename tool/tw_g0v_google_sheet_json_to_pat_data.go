package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	Url              string
	CaseId           string
	Gender           string
	RelationPlace    string
	Imported         string
	Age              string
	DateForHospital  string
	DateForConfirm   string
	Note             string
	DateForDeath     string
	DateForDischarge string
}

type DestinateFormat struct {
	Id            string `json:"id"`
	ReleaseDate   string `json:"リリース日"`
	MoveInDate    string `json:"移入日"`
	Place         string `json:"居住地"`
	RelationPlace string `json:"相關地點"`
	Imported      string `json:"境外或本土"`
	Age           string `json:"年代"`
	Gender        string `json:"性別"`
	IsExit        string `json:"退院"`
	Link          string `json:"link"`
}

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
	file, err := os.Open("./downloads/patients_raw.json")
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
			Url:             it2[keyToRowNumber["CDC 新聞稿 URL"]].(string),
			CaseId:          it2[keyToRowNumber["案例編號"]].(string),
			Gender:          it2[keyToRowNumber["性別"]].(string),
			RelationPlace:   it2[keyToRowNumber["相關地點"]].(string),
			Imported:        it2[keyToRowNumber["境外或是本土"]].(string),
			Age:             it2[keyToRowNumber["年齡層\n(~多歲)"]].(string),
			DateForHospital: it2[keyToRowNumber["就醫日"]].(string),
			DateForConfirm:  it2[keyToRowNumber["確診日"]].(string),
		}

		if len(it2) > keyToRowNumber["其他公布說明"] {
			if v, ok := it2[keyToRowNumber["其他公布說明"]].(string); ok {
				newItem.Note = v
			}
		}

		if len(it2) > keyToRowNumber["死亡日"] {
			if v, ok := it2[keyToRowNumber["死亡日"]].(string); ok {
				newItem.DateForDeath = v
			}
		}
		if len(it2) > keyToRowNumber["解除隔離日"] {
			if v, ok := it2[keyToRowNumber["解除隔離日"]].(string); ok {
				newItem.DateForDischarge = v
			}
		}

		srcList = append(srcList, newItem)
	}
	// fmt.Println(srcList)

	dstList := []DestinateFormat{}
	for _, item := range srcList {
		newItem := DestinateFormat{
			Id:            item.CaseId,
			ReleaseDate:   item.DateForConfirm,
			RelationPlace: item.RelationPlace,
			Imported:      item.Imported,
			Link:          item.Url,
		}
		if item.Age != "" {
			newItem.Age = item.Age + "代"
		}
		if item.Gender != "" {
			newItem.Gender = item.Gender + "性"
		}
		if item.DateForDeath != "" || item.DateForDischarge != "" {
			newItem.IsExit = "〇"
		}
		dstList = append(dstList, newItem)
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
	out["data"] = dstList

	outBytes, _ := json.MarshalIndent(out, "", "    ")
	fmt.Println(string(outBytes))

}
