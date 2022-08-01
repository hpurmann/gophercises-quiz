package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := flag.String("questions", "problems.csv", "filepath of questions csv file")
	flag.Parse()

	fileContent, err := os.ReadFile(*filename)

	if err != nil {
		fmt.Println("Error: cannot read file ", *filename, err)
	}

	reader := csv.NewReader(strings.NewReader(string(fileContent)))
	fields, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error marshalling csv", err)
	}

	var numCorrect int

	for index, questionAnswer := range fields {
		question, actualAnswer := questionAnswer[0], questionAnswer[1]
		fmt.Printf("Question %d: %s = ", index+1, question)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := scanner.Text()

		if strings.TrimSpace(answer) == actualAnswer {
			numCorrect += 1
		}
	}

	fmt.Printf("You scored %d out of %d.\n", numCorrect, len(fields))
}
