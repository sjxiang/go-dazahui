// 采集 模拟 POST 请求

package main

import "os"

func main() {
	if len(os.Args) != 2 {
		panic("参数不足")
	}

	words := os.Args[1]
	Query(words)
}


/*

	1. 抓包 
		网页 - 右键菜单 检查 -  network

		
	2. 代码生成 - HTTP 请求
		点击请求 - 右键菜单 copy - copy as cURL(bash)
		https://curlconverter.com/#go 


	3. 生成代码解读
		创建请求(url、请求头)
		发送请求
		读取响应


	4. 生成 request body


	5. 解析 response body
		JSON 转 struct
		
		https://oktools.net/json2go


	6. 完善代码


*/
