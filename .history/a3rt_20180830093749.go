package main

import (
	"fmt"
	"encoding/json"
	"net/url"
	"os"
	"net/http"
	"log"
	"io/ioutil"
)

const ENDPOINT = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"

type Results struct {
	Perplexity float64 `json:"perplexity"`
	Reply string `json:"reply"`
}

type Responce struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Results []Results `json:"results"`
}

var testText = "default"

func testFunc(text string) string {
	testText += text+"!!!!!!\n"
	return testText
}

func a3rt(query string) (*Responce, error) {
	apikey := os.Getenv("APIKEY")
	values := url.Values{}
	values.Add("apikey", apikey)
	values.Add("query", query)
	fmt.Println("aertの中\napikey = ", apikey)
	fmt.Println("a3rtの中\nquery = ", query)


	// http.PostForm(ENDPOINT, values)の中身
	// https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk?apikey=apikey&query=query
	res, err := http.PostForm(ENDPOINT, values)

	if err != nil {
		log.Fatal("*PostFrom*\n", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("*ReadAll*\n", err)
	}
	jsonBytes := ([]byte)(body)
	responce := new(Responce)
	err = json.Unmarshal(jsonBytes, responce)

	if err := json.Unmarshal(jsonBytes, responce); err != nil {
		log.Fatal("*json.Unmarshal*\n", err)
	}
	return responce, err
}