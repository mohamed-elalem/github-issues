package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	cli "./github"
)

var query = flag.String("q", "", "Search query for github api")
var label = flag.String("l", "", "The labels issues should contain")
var outputDirectory = flag.String("o", "./issues", "Output directory")

var usage = `Fetches github issues and formats the results as a simple html tables with some helpful information about each issue`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
		fmt.Println()
		flag.PrintDefaults()
	}
	flag.Parse()
	if *label == "" {
		log.Fatal("-l cannot be empty")
	}
	if *query == "" {
		log.Fatal("-q cannot be empty")
	}
	err := cli.Run(*query, *label, *outputDirectory)
	if err != nil {
		log.Fatal(err)
	}
}
