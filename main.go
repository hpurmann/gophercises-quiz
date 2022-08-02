package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func printScore(score, totalQuestions int) {
	fmt.Printf("\nYou scored %d out of %d.\n", score, totalQuestions)
	os.Exit(0)
}

func main() {
	filename := flag.String("questions", "problems.csv", "filepath of questions csv file")
	limit := flag.Int("limit", 30, "Time limit in seconds to complete the quiz")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		exit(fmt.Sprint("Error: cannot read file ", *filename, err))
	}

	reader := csv.NewReader(file)
	fields, err := reader.ReadAll()

	if err != nil {
		exit(fmt.Sprint("Error marshalling csv", err))
	}

	fmt.Printf("Welcome! Your time limit is %d seconds. The timer will start after hitting enter.\n", *limit)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()

	var numCorrect int

	time.AfterFunc(time.Second*time.Duration(*limit), func() {
		printScore(numCorrect, len(fields))
	})

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

	printScore(numCorrect, len(fields))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
