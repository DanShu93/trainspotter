package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/DanShu93/trainspotter/departure"
)

func main() {
	var duration, throttle, bufferTime int
	flag.IntVar(&duration, "duration", 3600, "This program runs for <throttle> seconds.")
	flag.IntVar(&throttle, "throttle", 60, "This program runs every <throttle> seconds.")
	flag.IntVar(&bufferTime, "buffer", 600, "This program warns you <buffer> seconds before your line arives.")

	flag.Parse()

	args := flag.Args()
	if len(args) != 5 {
		fmt.Println("Error invalid arguments. 5 arguments are needed <api_key> <origin> <destination> <transit_mode> <line_name>")
		os.Exit(1)
	}

	apiKey := args[0]
	origin := args[1]
	destination := args[2]
	transitMode := args[3]
	lineName := args[4]

	departure.Watch(duration, throttle, bufferTime, apiKey, origin, destination, transitMode, lineName)
}