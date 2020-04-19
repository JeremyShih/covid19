package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	// "net/url"
	"os"
)

// script:
// curl 'https://www.cdc.gov.tw/RSS/RssXml/Hh094B49-DRwe2RR4eFfrQ?type=1' > downloads/cdc_rss.xml
// go run tw_cdc_rss_to_news.go > ../data/news.json
// url: https://www.cdc.gov.tw/RSS/RssXml/Hh094B49-DRwe2RR4eFfrQ?type=1

// https://siongui.github.io/2015/02/27/go-parse-rss2/
type Rss2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate  string `xml:"channel>pubDate"`
	ItemList []Item `xml:"channel>item"`
}

type Item struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Updated  string        `xml:"updated"`
	Comments string        `xml:"comments"`
}

func main() {
	// fmt.Println("start")
	file, err := os.Open("./downloads/cdc_rss.xml")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	fileData, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := Rss2{}
	err = xml.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(data)
	list := []map[string]interface{}{}
	for _, item := range data.ItemList {
		list = append(list, map[string]interface{}{
			"date": item.Updated,
			"text": item.Title,
			"url":  item.Link,
		})
	}

	out := map[string]interface{}{}
	out["newsItems"] = list
	outBytes, _ := json.MarshalIndent(out, "", "    ")
	fmt.Println(string(outBytes))

}
