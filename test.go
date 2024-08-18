package main

import (
	"goInpy/net/http"
	"goInpy/pgbar"
	"goInpy/random"
	"time"
)

func main() {

	var Data = random.GenRandomString(160, "")

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
