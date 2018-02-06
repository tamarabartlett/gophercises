package main

import (
	"encoding/csv"
	"flag"
	"fmt"

	"os"
)

func main() {
	filePtr := flag.String("file", "data/problems.csv", "problems csv file")
	flag.Parse()

	correctCount := 0
	problems := readCSV(*filePtr)

	for _, problem := range problems {
		fmt.Println(problem.Question)
		var input string
		fmt.Scanln(&input)

		if input == problem.Answer {
			correctCount++
		}
	}
	fmt.Println("Total Number of Questions: ", len(problems))
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
