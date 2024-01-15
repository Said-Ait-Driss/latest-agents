package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"time"
)

type UserAgent struct {
	UA  string  `json:"ua"`
	Pct float64 `json:"pct"`
}

type APIInfo struct {
	About string `json:"about"`
	Terms string `json:"terms"`
	Data  []UserAgent `json:"data"`
}



func GetLatestList() string {
	resp, err := http.Get("https://www.useragents.me/api")

	if err != nil {
		fmt.Println(err)
	}

	body, errr := ioutil.ReadAll(resp.Body)

	if errr != nil {
		fmt.Println(errr)
	}

	stringOfAgents := string(body)
	return stringOfAgents
}

func turnToJson(textInput string) [] UserAgent {

	var apiInfo APIInfo

	 err := json.Unmarshal([]byte(textInput), &apiInfo )
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return apiInfo.Data
}

func saveIntoFile(filename string,content string ) error {
	toSave := []byte(content)

	return ioutil.WriteFile(filename,toSave,0644)
}

func generateFileName() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05 MST")
	fileName := fmt.Sprintf("data_%s.json", formattedTime)
	return fileName
}