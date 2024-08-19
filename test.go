package main

import (
	"fmt"
	"goInpy/net/http"
	"goInpy/pgbar"
	"io/ioutil"
	"strings"
	"time"
)

func main() {

	// var Target = "http://192.168.136.134:81/hello.aspx"
	var Target = "https://exchange.test2.com/owa"
	var BProxy = "http://127.0.0.1:8080"

	var Data = "__VIEWSTATE=%2FwEPDwUKMjA3NjE4MDczNmRk&__EVENTVALIDATION=%2FwEdAAONW96SHvVdMKC8pO3Ztp1R%2BCxUk8xX210xTeVYEF%2FTW834O%2FGfAV4V4n0wgFZHr3c%3D&TextArea1=123&Button1=GO"

	resp, err := http.Post(Target, BProxy, nil, strings.NewReader(Data), false)
	if err != nil {
		fmt.Println("err", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	// pgbar.InitDBar(Data)
	for i, _ := range Data {
		pgbar.Play(i, Data)
		time.Sleep(time.Millisecond * time.Duration(20))
	}

}

/*

	var Data = map[string]string{
		"__VIEWSTATE":       "%2FwEPDwUKMjA3NjE4MDczNmRk",
		"__EVENTVALIDATION": "%2FwEdAAONW96SHvVdMKC8pO3Ztp1R%2BCxUk8xX210xTeVYEF%2FTW834O%2FGfAV4V4n0wgFZHr3c%3D",
		"TextArea1":         "123",
		"Button1":           "GO",
	}

*/
