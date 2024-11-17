package solutions

import (
	"fmt"
	"encoding/csv"
	"os"
	"strings"
	"time"
)

func Sol1(csv_path string, time_limit int) int {
	fmt.Println(csv_path)
	var correct int
	correct = 0
	file, err := os.Open(csv_path)

	if (err != nil) {
		return 0
	}
	csv_r := csv.NewReader(file)
	lines, err := csv_r.ReadAll()

	if (err != nil) {
		return 0
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(time_limit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
			case <-timer.C:
				fmt.Println("\nOut of time\n")
				return correct
		case answer := <-answerCh:
				if answer == p.a {
					correct++
				}
		}
	}

	return correct
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
			ret[i] = problem{
				q: line[0],
				a: strings.TrimSpace(line[1]),
			}
	}
	return ret

}

type problem struct {
	q string
	a string

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)

}
