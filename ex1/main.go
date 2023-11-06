package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func recordTime(timer *time.Timer, quit chan bool, totalCap *int, totalCorrect *int) {
	fmt.Println("starting new timer")
mainLoop:
	for {
		select {
		case <-timer.C:
			fmt.Println("Reached time limit!")
			printStats(totalCorrect, totalCap)
			os.Exit(0)
		case <-quit:
			fmt.Println("Stopping time limit check")
			break mainLoop
		}
	}
}

func printStats(totalCap *int, totalCorrect *int) {
	fmt.Printf("You got %v of %v correct", *totalCap, *totalCorrect)
}

func main() {
	file, err := os.Open("problems.csv")
	printErr(err)
	lines, err := csv.NewReader(file).ReadAll()
	printErr(err)
	limitFlag := flag.Int("limit", 30, "The time limit per question")
	flag.Parse()
	fmt.Printf("Asking questions with time limit of %v seconds \n", *limitFlag)
	totalCap := len(lines)
	totalCorrect := 0
	for _, line := range lines {
		timer := time.NewTimer(time.Duration(*limitFlag) * time.Second)
		quit := make(chan bool)
		go recordTime(timer, quit, &totalCap, &totalCorrect)
		question, answer := line[0], line[1]
		print(question, " => ")
		printErr(err)
		var userAnswer string
		fmt.Scanf("%v", &userAnswer)
		if userAnswer == answer {
			fmt.Println("correct")
			totalCorrect += 1
		} else {
			fmt.Println("Incorrect")
		}
		quit <- true
	}
}
