package collect

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	url := ""

	reply, err := Query(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(*reply)
}