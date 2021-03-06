package departure

import (
	"time"
	"fmt"
)

func Watch(
	duration,
	throttle,
	bufferMin,
	bufferMax,
	offsetTime int,
	apiKey,
	origin,
	destination,
	transitMode string,
	lineNames []string,
	isWalk bool,
) {
	printStatus(bufferMin, bufferMax, offsetTime, apiKey, origin, destination, transitMode, lineNames, isWalk)

	ticker := time.NewTicker(time.Second * time.Duration(throttle))
	go func() {
		for range ticker.C {
			printStatus(bufferMin, bufferMax, offsetTime, apiKey, origin, destination, transitMode, lineNames, isWalk)
		}
	}()

	time.Sleep(time.Second * time.Duration(duration))
	ticker.Stop()

	fmt.Println("DONE")
}

func printStatus(
	bufferMin,
	bufferMax,
	offsetTime int,
	apiKey,
	origin,
	destination,
	transitMode string,
	lineNames []string,
	isWalk bool,
) {
	status := getStatus(bufferMin, bufferMax, offsetTime, apiKey, origin, destination, transitMode, lineNames, isWalk)
	fmt.Println(status)
}

func getStatus(bufferMin, bufferMax, offsetTime int, apiKey, origin, destination, transitMode string, lineNames []string, isWalk bool) string {
	desiredDepTime := time.Now().Add(time.Duration(offsetTime) * time.Second)
	depTime, err := GetDepartureTime(origin, destination, apiKey, transitMode, lineNames, desiredDepTime, isWalk)
	if err != nil {
		return fmt.Sprintf("ERROR %s", err)
	} else {
		until := time.Until(depTime)
		untilSeconds := int(until.Seconds()) - offsetTime

		if untilSeconds <= bufferMax {
			if untilSeconds < bufferMin {
				return fmt.Sprintf("HURRY %d", untilSeconds)
			} else {
				return fmt.Sprintf("GO %d", untilSeconds)
			}
		} else {
			return fmt.Sprintf("WAIT %d", untilSeconds)
		}
	}
}
