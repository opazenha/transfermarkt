package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func getPerformance(playerID int) {
	url := fmt.Sprintf("https://tmapi-alpha.transfermarkt.technology/player/%d/performance", playerID)
	body := fetch(url)

	var playerResponse Response
	if err := json.Unmarshal(body, &playerResponse); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	jsonFile := fmt.Sprintf("performance/performance_%d.json", playerID)
	err := ioutil.WriteFile(jsonFile, body, 0644)
	if err != nil {
		log.Fatalf("Failed to write JSON to file %s: %v", jsonFile, err)
	}

	if !playerResponse.Success {
		log.Fatalf("Failed to fetch data: %v", playerResponse.Message)
	}
}

func findStringSubmatchRegex(input string, pattern string) [][]string {
	re := regexp.MustCompile(pattern)
	return re.FindAllStringSubmatch(input, -1)
}
