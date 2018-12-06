package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() int {
	f, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	result := 0
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		result += i
	}

	return result
}

func part2() int {
	f, err := os.Open(`input.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// input の数値を配列にする
	ba, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	inputData := strings.Split(string(ba), "\n")

	result := 0
	resultMap := make(map[int]bool)

L:
	for {
		for _, d := range inputData {
			i, err := strconv.Atoi(d)
			if err != nil {
				continue
			}
			result += i
			if _, ok := resultMap[result]; ok {
				break L
			} else {
				resultMap[result] = true
			}
		}
	}

	return result
}
