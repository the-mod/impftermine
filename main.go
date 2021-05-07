package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	protocol      = "https"
	host          = "229-iz.impfterminservice.de"
	path          = "/rest/suche/termincheck"
	queryTemplate = "plz={{zipCode}}&leistungsmerkmale={{vaccines}}"
)

func createRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("name", "value")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "2350")
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	return req
}

func doRequest(url string) ([]byte, error) {
	req := createRequest(url)
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, err
}

func main() {
	zipCodes := [...]string{"70376", "76530"}
	vaccines := "L920,L921,L922,L923"

	resInit, _ := doRequest("https://www.impfterminservice.de")
	fmt.Println(resInit)

	for _, zipCode := range zipCodes {
		r := strings.NewReplacer(
			"{{zipCode}}", zipCode,
			"{{vaccines}}", vaccines)

		query := r.Replace(queryTemplate)
		url := fmt.Sprintf("%s://%s%s?%s", protocol, host, path, query)

		fmt.Println(url)

		res, err := doRequest(url)

		if err != nil {
			fmt.Println("error while requesting zip " + zipCode)
		}

		fmt.Println("Response: " + string(res))
	}
}
