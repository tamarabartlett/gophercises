package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"strconv"
	"time"

	"os"
)

func main() {
	filePtr := flag.String("file", "data/problems.csv", "problems csv file")
	timerPtr := flag.String("timer", "30", "timer for quiz")
	flag.Parse()

	timeInSeconds, err := strconv.Atoi(*timerPtr)
	if err != nil {
		fmt.Println("Invalid timer input. Defaulting to 30 seconds")
		timeInSeconds = 30
	}

	correctCount := 0
	problems := readCSV(*filePtr)

	timer := time.NewTimer(time.Second * time.Duration(timeInSeconds))
	defer timer.Stop()

	go func() {
		<-timer.C
		printTestResults(len(problems), correctCount)
		os.Exit(0)
	}()

	for _, problem := range problems {
		fmt.Println(problem.Question)
		var input string
		fmt.Scanln(&input)

		if input == problem.Answer {
			correctCount++
		}
	}
	printTestResults(len(problems), correctCount)
}

func printTestResults(numberOfProblems int, correctCount int) {
	fmt.Println("Total Number of Questions: ", numberOfProblems)
	fmt.Println("Correct: ", correctCount)
}

type problem struct {
	Question string
	Answer   string
}

func readCSV(csvString string) []problem {
	f, err := os.Open(csvString)
	if err != nil {
		fmt.Println("File name unreadable. Opening /data/problems.csv")
		f, _ = os.Open("data/problems.csv")
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var problems []problem
	for _, line := range lines {
		problems = append(problems, problem{
			Question: line[0],
			Answer:   line[1],
		})
	}
	return problems
}
