package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
)

func fetch(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")
	req.Header.Add("Accept-Language", "en,pt;q=0.9")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer response.Body.Close()

	// Handle gzipped response
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			log.Fatalf("Failed to create gzip reader: %v", err)
		}
		defer reader.Close()
	default:
		reader = response.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	return body
}
