package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	yandexEndpoint = "translate.yandex.net/api/v1.5/tr.json"
	translateText  = "translate"
	detectLanguage = "detect"
)

var (
	inputText string
	apiKey    string
)

type (
	request  *http.Request
	response struct {
		StatusCode     int      `json:"code"`
		Language       string   `json:"lang"`
		TranslatedText []string `json:"text"`
	}
)

func init() {
	apiKey = os.Getenv("KEY")
	flag.StringVar(&inputText, "text", "", "text to translate")
	flag.Parse()
}

func main() {
	if inputText == "" {
		flag.Usage()
		log.Fatal(fmt.Errorf("No text specified for translation").Error())
		return
	}

	url := buildRequestURL(yandexEndpoint, translateText, inputText, "en-es")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalf(fmt.Errorf("Unable to generate POST request").Error())
	}

	addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf(fmt.Errorf("Unsuccessful POST request").Error())
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	r := response{}
	decoder.Decode(&r)
	if r.StatusCode == 200 {
		fmt.Println(r.TranslatedText)
	}
}

func addHeaders(r *http.Request) {
	r.Header.Add("content-type", "multipart/form-data;")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "*/*")
	r.Header.Add("Host", "translate.yandex.net")
	r.Header.Add("Cache-Control", "no-cache")
}

func buildRequestURL(host, stub, text, lang string) string {
	return fmt.Sprintf("https://%s/%s?key=%s&text=%s&lang=%s", host, stub, apiKey, text, lang)
}
