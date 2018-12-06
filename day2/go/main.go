package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(part1())
}

func part1() int {
	f, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	twoLetterCount := 0
	threeLetterCount := 0
	var letterCounter map[byte]int
	for s.Scan() {
		letterCounter = map[byte]int{}
		for _, b := range s.Bytes() {
			letterCounter[b]++
		}

		twoLetterCounted := false
		threeLetterCounted := false
		for _, v := range letterCounter {
			if !twoLetterCounted && v == 2 {
				twoLetterCount++
				twoLetterCounted = true
			} else if !threeLetterCounted && v == 3 {
				threeLetterCount++
				threeLetterCounted = true
			}
		}
	}

	return twoLetterCount * threeLetterCount
}
