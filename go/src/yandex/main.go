package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func init() {
	apiKey = os.Getenv("YANDEX_KEY")
	flag.StringVar(&inputText, "text", "", "text to translate")
	flag.Parse()
}

func main() {
	if inputText == "" {
		flag.Usage()
		log.Fatal(fmt.Errorf("No text specified for translation").Error())
		return
	}

	url := buildRequestURL(translateText, inputText, translateDirection(English, Spanish))
	newReq, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalf(fmt.Errorf("Unable to generate POST request").Error())
	}

	req := request{
		HttpRequest: newReq,
	}
	req.addHeaders()
	res, err := http.DefaultClient.Do(req.HttpRequest)
	if err != nil {
		log.Fatalf(fmt.Errorf("Unsuccessful POST request").Error())
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	r := response{}
	decoder.Decode(&r)
	if r.StatusCode == 200 {
		fmt.Println(r.TranslatedText)
	} else {
		log.Fatalf(fmt.Errorf("Could not handle Response [%v]", r).Error())
	}
}

// func Gzip(h http.Handler) http.Handler {
// 	wrapper := func(w http.ResponseWriter, r *http.Request) {
// 		const coding = "gzip"
// 		if !strings.Contains(r.Header.Get("Accept-Encoding"), coding) {
// 			h.ServeHTTP(w, r)
// 			return
// 		}
//
// 		w.Header().Set("Content-Encoding", coding)
// 		w.Header().Set("Vary", "Accept-Encoding") // See http://www.fastly.com/blog/best-practices-for-using-the-vary-header/
// 		gz := gzWriterPool.Get().(*gzip.Writer)
// 		gz.Reset(w)
// 		defer func() {
// 			gz.Close()
// 			gzWriterPool.Put(gz)
// 		}()
//
// 		h.ServeHTTP(&gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
// 	}
//
// 	return http.HandlerFunc(wrapper)
// }

func (req request) addHeaders() {
	req.HttpRequest.Header.Add("content-type", "multipart/form-data;")
	req.HttpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.HttpRequest.Header.Add("Accept", "*/*")
	req.HttpRequest.Header.Add("Host", "translate.yandex.net")
	req.HttpRequest.Header.Add("Cache-Control", "no-cache")
}

func buildRequestURL(stub, text, lang string) string {
	parameters := url.Values{}
	parameters.Add("key", apiKey)
	parameters.Add("text", text)
	parameters.Add("lang", lang)

	return fmt.Sprintf("https://%s/%s?%s", yandexEndpoint, stub, parameters.Encode())
}
