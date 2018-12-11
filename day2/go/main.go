package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(part2())
}

func part1() int {
	f, err := os.Open("input.txt")
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

func part2() string {
	f, err := os.Open("input.txt.org")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// あらかじめ input を配列にする
	s := bufio.NewScanner(f)

	inputData := [][]byte{}
	for s.Scan() {
		copySrc := s.Bytes()
		copyDest := make([]byte, len(copySrc), cap(copySrc))
		copy(copyDest, copySrc)
		inputData = append(inputData, copyDest)
	}

	var result []byte
	for i, src := range inputData {
		for _, target := range inputData[i+1:] {
			tmp := dripMatchByte(src, target)
			if len(result) < len(tmp) {
				result = tmp
			}
		}
	}
	return string(result)
}

func dripMatchByte(l, r []byte) []byte {
	if len(l) != len(r) {
		log.Fatalf("len(l) != len(r): len(%d) != len(%d)", len(l), len(r))
	}
	var result []byte
	for i := 0; i < len(l); i++ {
		if l[i] == r[i] {
			result = append(result, l[i])
		}
	}
	return result
}
