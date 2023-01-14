
// 采集，模拟 POST 请求

package collect


import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


type Request struct {}

type Response struct {}


func Query(url string) (*Response, error) {

	client := &http.Client{}

	body:= Request{}

	data, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("accept", "application/json, text/plain, */*")
	
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()  

	// 读取响应
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	if resp.StatusCode != 200 {
		log.Fatalf("bad statusCode: %v", resp.StatusCode)
	}

	reply := &Response{}
	err = json.Unmarshal(bodyText, reply)
	if err != nil {
		return nil, err
	}

	return reply, err
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

