package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

func main() {
	strategy1()
}

type record struct {
	timestamp *time.Time
	message   string
}

func strategy1() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(b)

	lineCount := bytes.Count(b, []byte{'\n'})
	records := make([]record, lineCount, lineCount)
	for i := 0; i < lineCount; i++ {
		s, err := buf.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// TODO: Field 使わずにいい感じにパースする方法がありそう
		rawRecords := strings.Fields(strings.TrimRight(s, "\n"))
		t, err := time.Parse("[2006-01-02 15:04]", strings.Join(rawRecords[:2], " "))
		if err != nil {
			log.Fatal(err)
		}
		records[i] = record{
			timestamp: &t,
			message:   strings.Join(rawRecords[2:], " "),
		}
	}

	// chronological order
	sort.Slice(records, func(i, j int) bool { return records[i].timestamp.Before(*records[j].timestamp) })

	// debug
	for _, r := range records {
		fmt.Println(*r.timestamp, r.message)
	}

}
