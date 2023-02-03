package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)


func main() {
	sig, err := sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

/* 文件校验和值
	
$ cat sha1.go | sha1sum

	*/ 
func sha1Sum(fileName string ) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	var r io.Reader = file

	// io.CopyN(os.Stdout, r, 100)
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)  
	return fmt.Sprintf("%x", sig), nil
}


/* io

	io.Reader 
	io.Writer

*/