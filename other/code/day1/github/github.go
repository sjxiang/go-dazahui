package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	fmt.Println(githubInfo(ctx, "sjxiang"))
}


// githubInfo 返回 login 用户的姓名以及公开仓库数量
func githubInfo(ctx context.Context, login string) (string, int, error) {
	
	// 创建请求 
	url := "https://api.github.com/users/" + url.PathEscape(login)	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("user-agent", "Go")

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%#v - %s", url, resp.Status)
		/*		
			log.Printf("%#v - %s", url, resp.StatusCode)
			os.Exit(1)
		*/
	}

	// fmt.Printf("content-type: %s\n", resp.Header.Get("content-type"))
	
	// 读取响应
	var r struct {  // 匿名 struct
		Name        string `json:"login"`
		NumRepos    int    `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}


/*

$ curl -i -H 'User-Agent:go' https://api.github.com/users/sjxiang
	-i 可以显示响应 response header 信息
	-H 

	*/



/* io

	type Closer interface {
		Close() error
	}

	type Reader() interface {
		Read(p []byte) (n int, err error)
	}

	io.Copy(os.Stdout, src)  // 一大坨，乱糟糟的

	*/


/* JSON <-> Go

	true/false <-> true/false
	string     <-> string
	null       <-> nil
	number     <-> float64、float43、int8、...
	array      <-> []interface{}
	object     <-> map[string]interface{}、struct

	*/


/* JSON 编码/解码 API

	JSON -> io.Reader -> Go: json.Decoder
	JSON -> []byte    -> Go: json.Unmarshal

	Go -> io.Writer -> JSON: json.Encoder
	Go -> []byte    -> JSON: json.Marshal

*/