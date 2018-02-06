package main

import (
	"encoding/csv"
	"fmt"

	"os"
)

func main() {

	problems := readCSV("problems.csv")
	fmt.Println(problems)

}

type problem struct {
	Question string
	Answer   string
}

func readCSV(csvString string) []problem {

	f, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
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
