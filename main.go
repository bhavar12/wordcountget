package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ReqBody is type for input request
type ReqBody struct {
	Input string `json:"input"`
}

// Resp is type to provide the req resp
type Resp struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func main() {
	url := "http://localhost:8080/count"
	method := "POST"
	client := &http.Client{}
	obj := ReqBody{
		Input: "my name is chandrakant and me and my chandrakant and you and me working on golang prj and go lang prj is good and chandrakant is also good and his team is also good",
	}

	requestData, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("failed to marshal request input data: %v", err)
	}
	payload := strings.NewReader(string(requestData))
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		var resp []Resp
		err = json.Unmarshal(body, &resp)
		if err != nil {
			fmt.Printf("failed to unmarshal resp data: %+v ", err)
		} else {
			fmt.Printf("resp data %+v", resp)
		}
	}
}
