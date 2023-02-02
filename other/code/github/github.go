package main

import (
	"context"
	"encoding/json"
	"fmt"
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

	url := "https://api.github.com/users/" + url.PathEscape(login)
	
	// resp, err := http.Get(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}
	
	defer resp.Body.Close()
	// fmt.Printf("content-type: %s\n", resp.Header.Get("content-type"))
	// var r Reply

	var r struct {
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
	type Reply struct {
		Name        string `json:"login"`
		NumRepos    int    `json:"public_repos"`
	}
*/

/*

$ curl -i -H 'User-Agent:go' https://api.github.com/users/sjxiang
	-i 可以显示响应 response header 信息
	-H 

	*/



/* "io"

	type Closer interface {
		Close() error
	}

	type Reader() interface {
		Read(p []byte) (n int, err error)
	}

	io.Copy(os.Stdout, resp.Body)  // 一大坨，乱糟糟的
		
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