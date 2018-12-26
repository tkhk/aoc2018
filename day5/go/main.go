package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func generateCaseMap() map[byte]byte {
	m := map[byte]byte{}
	for _, l := range []byte("abcdefghijklmnopqrstuvwxyz") {
		u := bytes.ToUpper([]byte{l})[0]
		m[l] = u
		m[u] = l
	}
	return m
}

func main() {
	// part1()
	fmt.Println(part2())
}

var caseMap = map[byte]byte{}

func part1() int {
	caseMap = generateCaseMap()

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	before := b
	var result []byte
	for {
		after := deleteDuplicate(before)
		if len(before) == len(after) {
			result = after
			break
		}
		before = after
	}

	return len(result)
}

func deleteDuplicate(src []byte) []byte {
	deleteMarks := make([]int, len(src))
	markIndex := 0
	nextPassFlag := false
	for i := range src {
		if nextPassFlag {
			deleteMarks[markIndex] = i
			markIndex++
			nextPassFlag = false
			continue
		}
		if _, ok := caseMap[src[i]]; !ok {
			if src[i] == '\n' {
				continue
			}
			log.Fatalf("Does not exist in caseMap: Unexpected string: %x", src[i])
		}
		if caseMap[src[i]] == src[i+1] {
			deleteMarks[markIndex] = i
			markIndex++
			nextPassFlag = true
		}
	}

	if markIndex == 0 {
		return src
	}

	result := make([]byte, len(src)-markIndex)
	resultIndex := 0
	deleteIndex := 0
	for i := range src {
		if i == deleteMarks[deleteIndex] {
			deleteIndex++
			continue
		}
		result[resultIndex] = src[i]
		resultIndex++
	}
	return result
}

func part2() int {
	caseMap = generateCaseMap()
	ba, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	before := ba
	min := len(ba)
	for _, b := range []byte("abcdefghijklmnopqrstuvwxyz") {
		var result []byte
		for {
			after := deleteDuplicateAndByte(before, b)
			if len(before) == len(after) {
				result = after
				break
			}
			before = after
		}
	}
	return min
}

func deleteDuplicateAndByte(src []byte, removeByte byte) []byte {
	deleteMarks := make([]int, len(src))
	markIndex := 0
	nextPassFlag := false
	for i := range src {
		if src[i] == removeByte || src[i] == caseMap[removeByte] {
			deleteMarks[markIndex] = i
			markIndex++
			continue
		}
		if nextPassFlag {
			deleteMarks[markIndex] = i
			markIndex++
			nextPassFlag = false
			continue
		}
		if _, ok := caseMap[src[i]]; !ok {
			if src[i] == '\n' {
				continue
			}
			log.Fatalf("Does not exist in caseMap: Unexpected string: %x", src[i])
		}
		if caseMap[src[i]] == src[i+1] {
			deleteMarks[markIndex] = i
			markIndex++
			nextPassFlag = true
		}
	}

	if markIndex == 0 {
		return src
	}

	result := make([]byte, len(src)-markIndex)
	resultIndex := 0
	deleteIndex := 0
	for i := range src {
		if i == deleteMarks[deleteIndex] {
			deleteIndex++
			continue
		}
		result[resultIndex] = src[i]
		resultIndex++
	}
	return result
}
