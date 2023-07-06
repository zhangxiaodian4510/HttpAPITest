package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	rsps, err := http.Get("https://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=18618224510")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)

	if err != nil {
		fmt.Println("read error", err)
		return
	}

	fmt.Println(string(body))
}
