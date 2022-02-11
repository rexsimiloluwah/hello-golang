package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	q string
	a interface{}
}

func main() {
	timeLimit := flag.Int("int", 10, "Time limit for the quiz.")
	csvFileName := flag.String("csv", "problems.csv", "A csv file containing the problems in question,answer format.")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		errorExit(fmt.Sprintf("Failed to open file: %s.", *csvFileName))
	}

	// Reading the lines from the csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		errorExit(fmt.Sprintf("Failed to parse the csv file: %s.", err))
	}

	// timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// fmt.Println(lines)
	// fmt.Println(parseProblems(lines))

	// keep track of correct answers
	correctCount := 0
	// print out every question and answer
	for i, problem := range parseProblems(lines) {
		fmt.Printf("Problem #%d: %s\n", i, problem.q)
		// using go-routines for the timer part
		// create a channel for the answer, a channel basically provides a communication mechanism between two go-routines
		answerCh := make(chan string)
		// the go-routine
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("TIME OUT CHAPPIE!, You got %d out of %d answers correctly!", correctCount, len(parseProblems(lines)))
			return
		case answer := <-answerCh:
			if answer == problem.a {
				correctCount++
			}
		}
	}

	fmt.Printf("You got %d out of %d answers correctly!", correctCount, len(parseProblems(lines)))
}

func errorExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// parse the lines from the .csv file
func parseProblems(lines [][]string) []Problem {
	ret := make([]Problem, len(lines)-1) // len(lines)-1 to ignore the header
	for i, line := range lines {
		if i > 0 { // to ignore the header
			ret[i-1] = Problem{
				strings.TrimRight(line[0], " "),
				strings.TrimRight(line[1], " "),
			}
		}
	}

	return ret
}

