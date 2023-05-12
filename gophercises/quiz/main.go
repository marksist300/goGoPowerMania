package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var problemMap map[string]int = map[string]int{}

// Insert everything into the map the Q's as strings; answers as Numbers
func insertQuestions(probMap map[string]int, data []string) {
	var sum, answer string = data[0], data[1]
	var number int = parseNums(answer)
	probMap[sum] = number
}

// convert string number to int types
func parseNums(strNum string) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Printf("Error converting string %v to number\n", strNum)
		fmt.Println(err)
	}
	return num
}

// Take in all the data from the CSV and parse it.
func readInFileData(probMap map[string]int) {
	var csvFileName *string = flag.String("csv", "problems.csv", "Please provide a valid CSV file in the format: 'Question, Answer', where Answers are numeric (intger) values")

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvRead := csv.NewReader(file)
	for {
		data, err := csvRead.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		insertQuestions(probMap, data)
	}
	return
}

// runs the actual game loops through questions while the total score is lower than 10
func displayQuestions(probMap map[string]int, correctCount *int) {
	var timerDuration time.Duration = 10 * time.Second
	var timer *time.Timer = time.NewTimer(timerDuration)
	var stop = make(chan bool)
	go func() {
		<-timer.C
		fmt.Println("Time's up!")
		fmt.Printf("Your score was %v\n", *correctCount)
		close(stop)
		os.Exit(0)
	}()

	for k, answer := range probMap {
		var userVal int
		fmt.Printf("What does %v equal?\n", k)
		_, err := fmt.Scanln(&userVal)
		if err != nil {
			fmt.Println("Invalid Input, input must be a number", err)
			return
		}
		if userVal == answer {
			*correctCount++
		}

		if *correctCount >= 10 {
			stop <- true
			timer.Stop()
			fmt.Printf("Well done you got %v answers correct!\n", *correctCount)
			break
		}

		select {
		case <-stop:
			break
		default:
		}
	}
	return
}

func main() {
	var correctAnswers int
	readInFileData(problemMap)
	displayQuestions(problemMap, &correctAnswers)
}
