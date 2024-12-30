package main

import (
	"fmt"
	"os"
)

func previousMatches(playerID int) {
	url := fmt.Sprintf("https://www.transfermarkt.com/ceapi/previousMatches/player/%d?limit=300", playerID)
	body := fetch(url)

	f, err := os.Create(fmt.Sprintf("matches/previous_matches_%d.json", playerID))
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

func nextMatches(playerID int) {
	url := fmt.Sprintf("https://www.transfermarkt.com/ceapi/nextMatches/player/%d?limit=300", playerID)
	body := fetch(url)

	f, err := os.Create(fmt.Sprintf("matches/next_matches_%d.json", playerID))
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
