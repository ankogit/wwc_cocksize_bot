package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestGetGetNewCockSize(t *testing.T) {
	getNewCockSize()
}

func TestTestGen(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var all int
	var numbers []float64
	for i := 1; i < 50; i++ {
		numbers = append(numbers, float64(int(math.Round(test(float64(i), 1)))))
		all += int(math.Round(test(float64(i), 1)))
		fmt.Println(int(math.Round(test(float64(i), 1))))
	}

	fmt.Println("all", all/50)
	fmt.Println("median", CalcMedian(numbers))
	fmt.Println("---------")
}
