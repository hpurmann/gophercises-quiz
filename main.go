package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
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

	problems := parseLines(fields)

	fmt.Printf("Welcome! Your time limit is %d seconds. The timer will start after hitting enter.\n", *limit)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()

	var numCorrect int

	time.AfterFunc(time.Second*time.Duration(*limit), func() {
		printScore(numCorrect, len(fields))
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-c
		printScore(numCorrect, len(fields))
	}()

	for index, p := range problems {
		fmt.Printf("Question %d: %s = ", index+1, p.q)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := scanner.Text()

		if strings.TrimSpace(answer) == p.a {
			numCorrect += 1
		}
	}

	printScore(numCorrect, len(fields))
}

type problem struct {
	q string
	a string
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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
