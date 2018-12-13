package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(strategy2())
}

type record struct {
	timestamp *time.Time
	message   string
}

func strategy1() int {
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

	// search max asleep time guard
	asleepTimes := map[string]time.Duration{}
	var asleepStartTime *time.Time
	var asleepGuardID string
	for _, r := range records {
		messageWords := strings.Fields(r.message)
		switch messageWords[0] {
		case "Guard":
			asleepGuardID = messageWords[1]
		case "falls":
			asleepStartTime = r.timestamp
		case "wakes":
			if asleepStartTime == nil {
				log.Fatalf("asleepStartTime is nil: GuardID=%v", asleepGuardID)
			}
			asleepTimes[asleepGuardID] += ((r.timestamp.Sub(*asleepStartTime)) - time.Duration(1)*time.Minute)
			asleepStartTime = nil
		default:
			log.Fatalf("unexpected record prefix: %v", messageWords[0])
		}
	}
	var maxAsleepTimeGuardID string
	var tmpMaxAsleepTime time.Duration
	for k, v := range asleepTimes {
		if tmpMaxAsleepTime < v {
			maxAsleepTimeGuardID = k
			tmpMaxAsleepTime = v
		}
	}

	// TODO: 一番寝ている Guard 探す時のロジックのコピペ
	timeSlice := make([]int, 60, 60)
	var tmpGuardID string
	var tmpFallsMinute int
	for _, r := range records {
		messageWords := strings.Fields(r.message)
		switch {
		case messageWords[0] == "Guard":
			tmpGuardID = messageWords[1]
		case tmpGuardID == maxAsleepTimeGuardID && messageWords[0] == "falls":
			tmpFallsMinute = r.timestamp.Minute()
		case tmpGuardID == maxAsleepTimeGuardID && messageWords[0] == "wakes":
			for i := range timeSlice[tmpFallsMinute : r.timestamp.Minute()-1] {
				timeSlice[tmpFallsMinute+i]++
			}
		}
	}

	var maxMinuteIndex int
	maxMinute := 0
	for i, r := range timeSlice {
		if maxMinute < r {
			maxMinuteIndex = i
			maxMinute = r
		}
	}

	intGuardID, err := strconv.Atoi(strings.TrimLeft(maxAsleepTimeGuardID, "#"))
	if err != nil {
		log.Fatal(err)
	}
	return intGuardID * maxMinuteIndex
}

func strategy2() int {
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

	frequentlyMap := map[string][60]int{}
	var tmpGuardID string
	var tmpFallsMinute int
	for _, r := range records {
		messageWords := strings.Fields(r.message)
		switch messageWords[0] {
		case "Guard":
			tmpGuardID = messageWords[1]
		case "falls":
			tmpFallsMinute = r.timestamp.Minute()
		case "wakes":
			f := frequentlyMap[tmpGuardID]
			for i := 0; i < r.timestamp.Minute()-tmpFallsMinute; i++ {
				f[tmpFallsMinute+i]++
			}
			frequentlyMap[tmpGuardID] = f
		default:
			log.Fatalf("Unexpected messageWords[0]: %v", messageWords[0])
		}
	}

	tmpMax := 0
	tmpMaxMinute := 0
	var tmpMaxGuardID string
	for k, vs := range frequentlyMap {
		for i, v := range vs {
			if tmpMax < v {
				tmpMax = v
				tmpMaxGuardID = k
				tmpMaxMinute = i
			}
		}
	}

	intGuardID, err := strconv.Atoi(strings.TrimLeft(tmpMaxGuardID, "#"))
	if err != nil {
		log.Fatal(err)
	}
	return intGuardID * tmpMaxMinute
}
