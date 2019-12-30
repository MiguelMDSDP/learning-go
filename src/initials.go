package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]

	statistics := getStatistics(words)

	myPrint(statistics)
}

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		counter, found := statistics[initial]
		if found {
			statistics[initial] = counter + 1
		} else {
			statistics[initial] = 1
		}
	}

	return statistics
}

func myPrint(statistics map[string]int) {
	fmt.Println("**Word count starting with each letter:")

	for initial, counter := range statistics {
		fmt.Printf("%s = %d\n", initial, counter)
	}
}
