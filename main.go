package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type DictRequest struct {
	Source []string `json:"source"`
	TransType string `json:"trans_type"`
	RequestID string `json:"request_id"`
	Media string `json:"media"`
	OsType string `json:"os_type"`
	Dict bool `json:"dict"`
	Cached bool `json:"cached"`
	Replaced bool `json:"replaced"`
	Detect bool `json:"detect"`
	BrowserID string `json:"browser_id"`
}


type DictResponse struct {
	Target []string `json:"target"`
	Rc int `json:"rc"`
	Confidence float64 `json:"confidence"`
}
// type DictResponse struct {
// 	Rc int `json:"rc"`
// 	Wiki struct {
// 	} `json:"wiki"`
// 	Dictionary struct {
// 		Prons struct {
// 			EnUs string `json:"en-us"`
// 			En string `json:"en"`
// 		} `json:"prons"`
// 		Explanations []string `json:"explanations"`
// 		Synonym []string `json:"synonym"`
// 		Antonym []string `json:"antonym"`
// 		WqxExample [][]string `json:"wqx_example"`
// 		Entry string `json:"entry"`
// 		Type string `json:"type"`
// 		Related []interface{} `json:"related"`
// 		Source string `json:"source"`
// 	} `json:"dictionary"`
// }


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `缺少参数`)
		os.Exit(1)
	}
	
	word := os.Args[1]
	query(word)
}

/*
client := &http.Client{}
	var data = 
	strings.NewReader(`{"source":["japan",""],"trans_type":"auto2zh","request_id":"web_fanyi","media":"text","os_type":"web","dict":true,"cached":true,"replaced":true,"detect":true,"browser_id":"09bd226c7c99858d9978a48513b38ee0"}`)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/translator", data)
	

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

*/

func query(word string) {

	client := &http.Client{}

	requestBody:= DictRequest{
		TransType: "auto2zh",
		Source: []string{word},
	}
	var data = strings.NewReader(`{"source":["japan",""],"trans_type":"auto2zh","request_id":"web_fanyi","media":"text","os_type":"web","dict":true,"cached":true,"replaced":true,"detect":true,"browser_id":"09bd226c7c99858d9978a48513b38ee0"}`)
	
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	/*
		io.Reader & stream 数据流

			网络 io 与浏览器 v8 引擎渲染速率不匹配，需要将就下，读一点处理一点，需要多次调用
			满足流式数据传输需求，多次调用

		
		struct {
			src  // 数据源（初始化字符串）
			cur  // 当前下标（已读计数）
		}
	
		read(p []byte) (n int, err error)  // p 申请内存，读取数据往里面塞；n 偏移量

	*/

	// 创建请求
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/translator", bytes.NewBuffer(requestData))
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("app-name", "xy")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("t-authorization", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJicm93c2VyX2lkIjoiMDliZDIyNmM3Yzk5ODU4ZDk5NzhhNDg1MTNiMzhlZTAiLCJpcF9hZGRyZXNzIjoiMTEyLjgwLjE1Ni4xMCIsInRva2VuIjoicWdlbXY0anIxeTM4anlxNnZodmkiLCJ2ZXJzaW9uIjoxLCJleHAiOjE2NzM3MTEwNzl9.u5dbfZVMsLIlhUKnctGbrYuhix9dzmJDnyf8MA5lVEQ")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()  

	// 读取响应
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyText))
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}

	reply := &DictResponse{}
	err = json.Unmarshal(bodyText, reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.Target) // BYTE 数组 转 string

}



/*

1. 抓包 
	网页 - 右键菜单 检查 -  network

2. 代码生成 - HTTP 请求
	点击请求 - 右键菜单 copy - copy as cURL(bash)
	https://curlconverter.com/#go 

3. 生成代码解读
	创建请求
	设置请求头
	发送请求
	读取响应

4. 生成 request body

5. 解析 response body
	JSON 转 struct
	
	https://oktools.net/json2go

6. 完善代码


*/

