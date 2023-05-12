package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	file, err := os.Open("problems.csv")
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
}

func displayQuestions(probMap map[string]int, correctCount *int) {
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
			fmt.Printf("Well done you got %v answers correct!", *correctCount)
			break
		}
	}
}

func main() {
	var correctAnswers int
	readInFileData(problemMap)
	displayQuestions(problemMap, &correctAnswers)
	fmt.Println("*** End of File ***")
}
