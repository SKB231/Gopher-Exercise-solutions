package main

import (
	"fmt"
	"os"
	"strings"
)

func readData() []string {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
	}
	dataArr := make([]byte, 2)
	data := make([]byte, 0)
	for {
		n, err := file.Read(dataArr)
		if err != nil {
			break
		}
		data = append(data, dataArr[:n]...)
	}
	questions := strings.Split(string(data), "\n")
	return questions
}

func main() {
	questions := readData()
	numCorrect := 0
	numQuestions := 0
	for _, question := range questions {
		if len(question) <= 0 {
			break
		}
		res := (strings.Split(question, ","))
		q, ans := res[0], res[1]
		var userAnswer string
		numQuestions += 1
		fmt.Println(q)
		fmt.Scan(&userAnswer)
		if userAnswer == ans {
			fmt.Println("correct answer")
			numCorrect += 1
		} else {
			fmt.Println("Wrong answer")
		}
	}
	fmt.Fprintf(os.Stdout, `You got %v questions correct out of %v`, numCorrect, numQuestions)
}
