package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var problemMap map[string]string = map[string]string{}

// Insert everything into the map the Q's as strings; answers as Numbers
func insertQuestions(probMap map[string]string, data []string) {
	var sum, answer string = data[0], data[1]
	// var number int = parseNums(answer)
	probMap[sum] = answer
}

// convert string number to int types
// func parseNums(strNum string) int {
// 	num, err := strconv.Atoi(strNum)
// 	if err != nil {
// 		fmt.Printf("Error converting string %v to number\n", strNum)
// 		fmt.Println(err)
// 	}
// 	return num
// }

func exit(msg string, code int) {
	fmt.Println((msg))
	os.Exit(code)
}

// Take in all the data from the CSV and parse it.
func readInFileData(probMap map[string]string) {
	var csvFileName *string = flag.String("csv", "problems.csv", "Please provide a valid CSV file in the format: 'Question, Answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		// log.Fatal(err)
		exit(fmt.Sprintf("Failed to open file: %v", *csvFileName), 1)
	}
	defer file.Close()
	csvRead := csv.NewReader(file)
	for {
		data, err := csvRead.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			exit("Data could not be parsed from the fild", 1)
		}
		insertQuestions(probMap, data)
	}
	return
}

// runs the actual game loops through questions while the total score is lower than 10
func displayQuestions(probMap map[string]string, correctCount, questionsAttempted *int) {
	var timerDuration time.Duration = 10 * time.Second
	var timer *time.Timer = time.NewTimer(timerDuration)
	var stop = make(chan bool)
	go func() {
		<-timer.C
		fmt.Println("Time's up!")
		close(stop)
		exit(fmt.Sprintf("You score was %d out of %d correct", *correctCount, *questionsAttempted), 0)
	}()

	for k, answer := range probMap {
		var userVal string
		fmt.Printf("What does %v equal?\n", k)
		_, err := fmt.Scanln(&userVal)
		*questionsAttempted++
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
			break
		}

		select {
		case <-stop:
			break
		default:
		}
	}
	exit(fmt.Sprintf("You got %d out of %d correct", *correctCount, *questionsAttempted), 0)
}

func main() {
	var correctAnswers int
	var questionsAttempted int
	readInFileData(problemMap)
	displayQuestions(problemMap, &correctAnswers, &questionsAttempted)
}
