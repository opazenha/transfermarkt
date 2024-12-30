package main

import (
	"fmt"
	"os"
)

func getTransferHistory(playerID int) {
	url := fmt.Sprintf("https://www.transfermarkt.com/ceapi/transferHistory/list/%d", playerID)
	body := fetch(url)

	f, err := os.Create(fmt.Sprintf("transfers/transfer_history_%d.json", playerID))
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
