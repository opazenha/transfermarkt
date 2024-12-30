package main

import "time"

func main() {

}

func getPLayersData() {
	playerIDs := []int{193925, 694926, 258624, 518676, 797734, 351897, 629860, 614412, 153568, 949337, 686453, 594546, 779864, 754789, 654133, 1043833, 62200}

	for _, playerID := range playerIDs {
		getPerformance(playerID)
		previousMatches(playerID)
		nextMatches(playerID)
		getTransferHistory(playerID)
		getMarketValue(playerID)

		time.Sleep(5 * time.Second)
	}
}
