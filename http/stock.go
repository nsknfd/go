package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const appcode = "c030ba4543b54e36ae9f8eda519c8353"

func Get(rawurl string, params url.Values) ([]byte, error) {
	URL, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	URL.RawQuery = params.Encode()
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "APPCODE "+appcode)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// QueryKLine 查询K线图
func QueryKLine(beginDay string, code string, time string, fqType string) {
	rawurl := "http://stock.market.alicloudapi.com/realtime-k"
	params := url.Values{}

	params.Set("beginDay", beginDay)
	params.Set("code", code)
	params.Set("time", time)
	params.Set("type", fqType)

	response, err := Get(rawurl, params)
	if err != nil {
		fmt.Printf("请求失败:\r\n%v", err)
	} else {
		var resp map[string]interface{}
		json.Unmarshal(response, &resp)
		//fmt.Println(response)
		respBody := resp["showapi_res_body"].(map[string]interface{})
		dataList := respBody["dataList"].([]interface{})
		fmt.Println(dataList)
	}
}
