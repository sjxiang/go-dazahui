package main

import (
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

/*

解压缩
$ unzip -d ./tmp taxi-sha256.zip


计算 /tmp 下文件的 sha256 签名，看看它们是否与索引文件中的签名相符。

  - 打印已处理文件的数量
  - 如果有不匹配，打印出违规的文件，并以非零值退出程序。

*/


func fileSig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))  
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}


// 解析签名文件，返回 map (文件名-> sha256 签名 )
func parseSigFile(r io.Reader) (map[string]string, error) {
	sigs := make(map[string]string, 10)
	i := 1

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// 切开空白符间隔的字符串
		fields := strings.Fields(scanner.Text()) 
		if len(fields) != 2 {
			return nil, fmt.Errorf("line %d - bad line: %q", i, scanner.Text())
		}
		sigs[fields[1]] = fields[0]
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}


func main() {
	rootDir := "./tmp"
	file, err := os.Open(path.Join(rootDir, "sha256sum.txt"))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	sigs, err := parseSigFile(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	start := time.Now()
	ok := true 
	ch := make(chan result)
	
	for name, signature := range sigs {
		fileName := path.Join(rootDir, name) + ".bz2"
		go sigWoker(fileName, signature, ch)
	}

	for range sigs {
		r := <-ch

		if r.err != nil {
			fmt.Fprintf(os.Stderr, "error: %s - %s\n", r.fileName, err)
			ok = false
			continue
		}

		if !r.match {
			ok = false
			fmt.Printf("error: %s mismatch\n", r.fileName)
		}
	}


	duration := time.Since(start)
	fmt.Printf("处理 %d 个文件，耗时 %v \n", len(sigs), duration)

	if !ok {
		os.Exit(1)
	}
	
}


func sigWoker(fileName, signature string, ch chan result) {
	r := result{fileName: fileName}
	sig, err := fileSig(fileName)
	if err != nil {
		r.err = err
	} else {
		r.match = sig == signature
	}

	ch <- r
}



type result struct {
	fileName string
	err      error
	match    bool
}


