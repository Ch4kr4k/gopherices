package main

import (
	"flag"
	"fmt"
	"gopherices/solutions"
)

func main() {

	var csv_path string // variable to store csv

	fmt.Println("Main Func started")
	// argument parsing and pointer
	csv_file := flag.String("csv", "file.csv", "Path to csv (use -c for short)")
	time_limit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.StringVar(csv_file, "c", "file.csv", "Path to the CSV file (short flag)")
	flag.Parse()

	csv_path = *csv_file
	fmt.Println("correct answers :", solutions.Sol1(csv_path, *time_limit))
}
