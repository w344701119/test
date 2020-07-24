package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var DateTimeLayoutOne = "2006-01-02 15:04:05"
var DateTimeLayoutTwo = "2006-01-02T15:04:05"

func main() {
	param := os.Args
	var inputTime string
	if len(param) > 2 {
		inputTime = strings.Trim(param[1], " ")
		fmt.Println("input time:", inputTime)
	}

	var err error
	var reqClient *http.Client
	var req *http.Request
	var startTime = StringDateToTime("2020-03-17 00:00:00", DateTimeLayoutOne)
	if inputTime != "" {
		startTime = StringDateToTime(inputTime, DateTimeLayoutOne)
	}

	var es_host = "http://127.0.0.1:9200"
	var reqUrl = es_host + "/db_nlp3/_search"
	postMap := map[string]interface{}{
		"_source": []string{"source"},
		"size":    10000,
	}
	var saveFile = "./chart2020.txt"
	var sf *os.File
	sf, err = os.OpenFile(saveFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for true {
		endTime := startTime.Add(time.Second * 60)
		createTime := map[string]interface{}{
			"gte": startTime.Format(DateTimeLayoutTwo),
			"lte": endTime.Format(DateTimeLayoutTwo),
		}
		rangeData := map[string]interface{}{
			"create_time": createTime,
		}
		postMap["query"] = map[string]interface{}{"range": rangeData}
		jsonData, err := json.Marshal(postMap)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		reqClient = http.DefaultClient
		req, err = http.NewRequest(http.MethodGet, reqUrl, bytes.NewReader(jsonData))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		dataLen := len(jsonData)
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		req.Header.Set("Content-Length", strconv.Itoa(dataLen))
		var result *http.Response
		result, err = reqClient.Do(req)
		if err != nil {
			fmt.Println("request err:", err)
			os.Exit(0)
		}
		if result.StatusCode != 200 {
			fmt.Println("fail to request with err code:", result.StatusCode)
			os.Exit(0)
		}
		var reqData []byte
		defer result.Body.Close()
		reqData, err = ioutil.ReadAll(result.Body)
		if err != nil {
			fmt.Println("fail to read result")
			os.Exit(0)

		}
		res, err := simplejson.NewJson(reqData)
		if err != nil {
			fmt.Println("simple err:", err)
			os.Exit(0)
		}
		rows, err := res.Get("hits").Get("hits").Array()
		if err != nil {
			fmt.Println("get array  err:", err)
			os.Exit(0)
		}
		if len(rows) == 0 {
			startTime = startTime.Add(time.Second * 60)
			fmt.Println("startTime:", startTime.Format(DateTimeLayoutOne))
			fmt.Println("endTime:", endTime.Format(DateTimeLayoutOne))
			continue
		}
		for _, row := range rows {
			if rowMap, ok := row.(map[string]interface{}); ok {
				tmp := rowMap["_source"]
				if source, ok := tmp.(map[string]interface{}); ok {
					str := fmt.Sprint(source["source"])
					sf.WriteString(str + "\n")
				} else {
					fmt.Println(ok)
				}
			} else {
				fmt.Println(ok)
			}
		}
		fmt.Println("startTime:", startTime.Format(DateTimeLayoutOne))
		fmt.Println("endTime:", endTime.Format(DateTimeLayoutOne))
		startTime = startTime.Add(time.Second * 60)
	}
	err = sf.Close()

}

//datetime 类型转换为time
func StringDateToTime(dateTime string, timeLayout string) (theTime time.Time) {
	//重要：获取时区
	loc, _ := time.LoadLocation("Local")
	theTime, _ = time.ParseInLocation(timeLayout, dateTime, loc)
	return
}
