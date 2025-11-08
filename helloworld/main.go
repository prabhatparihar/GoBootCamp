package main

import (
	"fmt"
)

func main() {
	var r rating
	r = reviewSystem(r)
	err := storeReviewInCSV(r)
	if err != nil {
		fmt.Println("Getting error while storing Data in CSV")
	}
	printStarsAfterFeedBack(r)

}
