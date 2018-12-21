package main

import (
	"bytes"
	//	"fmt"
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
	part1()
}

var caseMap = map[byte]byte{}

func part1() {
	caseMap = generateCaseMap()

	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var targets []byte
	targets = b
	// for {
	duplicatedIndexes := searchDuplicatedIndexes(targets)
	// if len(duplicatedIndexes) == 0 {
	// 	break
	// }
	newTargets := make([]byte, len(targets)-len(duplicatedIndexes))

	newTargetsIndex := 0
	duplicatedIndexesIndex := 0
	for i := range newTargets {
		if i != duplicatedIndexes[duplicatedIndexesIndex] {
			continue
		}
		newTargets[i] = targets[i]
	}
	//}
}

func searchDuplicatedIndexes(b []byte) []int {
	removeIndexes := []int{}
	for i := 0; i < len(b); i++ {
		if _, ok := caseMap[b[i]]; !ok {
			if b[i] == '\n' {
				continue
			}
			log.Fatalf("Does not exist in caseMap: Unexpected string: %v", b[i])
		}

		if caseMap[b[i]] == b[i+1] {
			removeIndexes = append(removeIndexes, []int{i, i + 1}...)
			i++
		}
	}
	return removeIndexes
}
