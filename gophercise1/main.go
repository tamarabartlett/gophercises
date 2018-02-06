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
	fileName := csvString
	if fileName == "" {
		fileName = "data/problems.csv"
	}

	f, err := os.Open(fileName)
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
