package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type responseT struct {
	Request string `json:"request"`
	Result  int    `json:"result"`
	Traceid string `json:"traceid"`
}

func GetTestResult(url string) ([]responseT, float64) {
	var bodyArray []responseT
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Response Code: ", resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error Reading Body:", err)
		} else {
			bodyStr := string(body)
			fmt.Println("Body: ", bodyStr)
			err := json.Unmarshal([]byte(bodyStr), &bodyArray)
			if err != nil {
				fmt.Println("Error Parsing response:", err)
			}
		}
	}
	return bodyArray, duration.Seconds()
}
