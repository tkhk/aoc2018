package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() int {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(b)

	// input.txt を slice としてあらかじめ読み込む
	lineCount := bytes.Count(b, []byte{'\n'})
	inputData := make([]string, lineCount, lineCount)
	inputDatai := 0
	for {
		s, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		inputData[inputDatai] = strings.TrimRight(s, "\n")
		inputDatai++
	}

	// fabricMap 初期化
	fabricMap := make([][]string, 1000, 1000)
	for i := range fabricMap {
		fabricMap[i] = make([]string, 1000, 1000)
	}

	//
	for _, d := range inputData {
		rectData := strings.Fields(d)

		id := rectData[0]
		position := strings.Split(strings.TrimRight(rectData[2], ":"), ",")
		scale := strings.Split(rectData[3], "x")

		targetXPosition, err := strconv.Atoi(position[0])
		if err != nil {
			log.Fatal(err)
		}
		targetYPosition, err := strconv.Atoi(position[1])
		if err != nil {
			log.Fatal(err)
		}
		targetXScale, err := strconv.Atoi(scale[0])
		if err != nil {
			log.Fatal(err)
		}
		targetYScale, err := strconv.Atoi(scale[1])
		if err != nil {
			log.Fatal(err)
		}

		for x, columns := range fabricMap[targetXPosition : targetXPosition+targetXScale] {
			for y := range columns[targetYPosition : targetYPosition+targetYScale] {
				if len(fabricMap[targetXPosition+x][targetYPosition+y]) != 0 {
					fabricMap[targetXPosition+x][targetYPosition+y] = "X"
				} else {
					fabricMap[targetXPosition+x][targetYPosition+y] = id
				}
			}
		}
	}

	// X カウント
	count := 0
	for _, columns := range fabricMap {
		for _, row := range columns {
			if row == "X" {
				count++
			}
		}
	}
	return count
}

func part2() string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(b)

	// input.txt を slice としてあらかじめ読み込む
	lineCount := bytes.Count(b, []byte{'\n'})
	inputData := make([]string, lineCount, lineCount)
	inputDatai := 0
	for {
		s, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		inputData[inputDatai] = strings.TrimRight(s, "\n")
		inputDatai++
	}

	// fabricMap 初期化
	fabricMap := make([][]string, 1000, 1000)
	for i := range fabricMap {
		fabricMap[i] = make([]string, 1000, 1000)
	}

	//
	overlapFlagMap := map[string]bool{}
	for _, d := range inputData {
		rectData := strings.Fields(d)

		id := rectData[0]
		position := strings.Split(strings.TrimRight(rectData[2], ":"), ",")
		scale := strings.Split(rectData[3], "x")

		targetXPosition, err := strconv.Atoi(position[0])
		if err != nil {
			log.Fatal(err)
		}
		targetYPosition, err := strconv.Atoi(position[1])
		if err != nil {
			log.Fatal(err)
		}
		targetXScale, err := strconv.Atoi(scale[0])
		if err != nil {
			log.Fatal(err)
		}
		targetYScale, err := strconv.Atoi(scale[1])
		if err != nil {
			log.Fatal(err)
		}

		tmpOverlapFlag := false
		for x, columns := range fabricMap[targetXPosition : targetXPosition+targetXScale] {
			for y := range columns[targetYPosition : targetYPosition+targetYScale] {
				if len(fabricMap[targetXPosition+x][targetYPosition+y]) != 0 {
					overlapFlagMap[fabricMap[targetXPosition+x][targetYPosition+y]] = true
					tmpOverlapFlag = true
					fabricMap[targetXPosition+x][targetYPosition+y] = "X"

				} else {
					fabricMap[targetXPosition+x][targetYPosition+y] = id
				}
			}
		}

		overlapFlagMap[id] = tmpOverlapFlag
	}

	// Check overlap
	var result string
	for k, v := range overlapFlagMap {
		if !v {
			result = k
			break
		}
	}

	return result
}
