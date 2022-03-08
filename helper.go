package main

import (
	"math/rand"
	"strconv"
	"time"
)

func getNewCockSize() int {
	min := 1
	max := 50

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func getCockSizeMessage(cocksize int) string {
	emoji := "😭"

	if cocksize > 1 && cocksize < 5 {
		emoji = "😰"
	} else if cocksize >= 5 && cocksize < 10 {
		emoji = "😥"
	} else if cocksize >= 10 && cocksize < 15 {
		emoji = "😓"
	} else if cocksize >= 15 && cocksize < 20 {
		emoji = "😏"
	} else if cocksize >= 20 && cocksize < 30 {
		emoji = "😏"
	} else if cocksize >= 30 && cocksize < 40 {
		emoji = "🤤"
	} else if cocksize >= 40 && cocksize < 50 {
		emoji = "🤥"
	}

	return "My cock size is " + strconv.Itoa(cocksize) + "cm " + emoji
}
