// 采集，模拟 POST 请求

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


const huoshanTranslateURL = "https://translate.volcengine.com/web/dict/detail/v1/?msToken=&X-Bogus=DFSzswVLQDccSyCWSZ4BtfNSwbzy&_signature=_02B4Z6wo00001Voyv5wAAIDCD7A0ABNiQf1aMrsAADVN75f0tzy88XdILjrP3zz3BWO7yCXm6JelSVyl0saHffnfrBmRnPB96khcgciMB.tTtIpliBdiU.ZrJctsPq9ruUuFPOfP7ojme-6Ja1"


func Query(word string) {

	client := &http.Client{}

	body:= Request {
		Source: "youdao",
		Words: []string{word},
		SourceLanguage: "en",
		TargetLanguage: "zh",
		
	}


	data, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", huoshanTranslateURL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "i18next=zh-CN; ttcid=17287189e4014ac6ba87bdfa6c04052712; x-jupiter-uuid=16738753375446688; s_v_web_id=verify_lcyu5omb_OLOsChcZ_GmDh_4vIw_8vob_Gl5hp5wKx7vx; tt_scid=Dy5zxuOoFaztJdwNZJibfPbDGbFw.GJWzBHq1N8zHvfEvHwu7fEw7Uxy4H-DIFi32b99")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://translate.volcengine.com/mobile?category=&home_language=zh&ref=toolsdar.com&source_language=detect&target_language=zh&text=fuck")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="99", "Microsoft Edge";v="109", "Chromium";v="109"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.52")
	

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
	
	if resp.StatusCode != 200 {
		log.Fatalf("bad statusCode: %v", resp.StatusCode)
	}

	reply := &Response{}
	err = json.Unmarshal(bodyText, reply)
	if err != nil {
		log.Fatal(err)
	}

	
	fmt.Println(reply.Details)
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


/*


	strings.NewReader 或者 bytes.NewBuffer()，皆可

	io.Reader & stream 数据流

		网络 io 与浏览器 v8 引擎渲染速率不匹配，需要将就下，读一点处理一点，需要多次调用
		满足流式数据传输需求，多次调用

	
	struct {
		src  // 数据源（初始化字符串）
		cur  // 当前下标（已读计数）
	}

	read(p []byte) (n int, err error)  // p 申请内存，读取数据往里面塞；n 偏移量


	*/




