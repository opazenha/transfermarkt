package main

import (
	"fmt"
	"os"
)

func getMarketValue(playerID int) {
	url := fmt.Sprintf("https://www.transfermarkt.com/ceapi/marketValueDevelopment/graph/%d", playerID)
	body := fetch(url)

	f, err := os.Create(fmt.Sprintf("market_value/market_value_%d.json", playerID))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
