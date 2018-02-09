package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"os"
)

type problem struct {
	Question string
	Answer   string
}

func main() {
	filePtr := flag.String("file", "data/problems.csv", "problems csv file")
	timerPtr := flag.String("timer", "30", "timer for quiz")
	shufflePtr := flag.String("shuffle", "n", "shuffle the quiz questions: please type n or y")
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

	if *shufflePtr == "y" {
		shuffleProblems(problems, &correctCount)
	} else {
		problemsInOrder(problems, &correctCount)
	}

	printTestResults(len(problems), correctCount)
}

func readAnswerAndIncrementCount(p problem, correctCount *int) {
	fmt.Println(p.Question)
	var input string
	fmt.Scanln(&input)

	cleanInput := strings.ToLower(strings.Trim(input, " ,."))

	if cleanInput == strings.ToLower(p.Answer) {
		*correctCount++
	}
}

func problemsInOrder(problems []problem, correctCount *int) {
	for _, p := range problems {
		readAnswerAndIncrementCount(p, correctCount)
	}
}

func shuffleProblems(problems []problem, correctCount *int) {
	for len(problems) > 0 {
		randProblem := rand.Intn(len(problems))
		currentProblem := problems[randProblem]
		problems = append(problems[:randProblem], problems[randProblem+1:]...)

		readAnswerAndIncrementCount(currentProblem, correctCount)
	}
}

func printTestResults(numberOfProblems int, correctCount int) {
	fmt.Println("Total Number of Questions: ", numberOfProblems)
	fmt.Println("Correct: ", correctCount)
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
