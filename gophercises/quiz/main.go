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

func insertQuestions(probMap map[string]int, data []string) {
	var sum, answer string = data[0], data[1]
	num, err := strconv.Atoi(answer)
	if err == nil {
		probMap[sum] = num
	} else {
		fmt.Println("Error converting answer to number", err)
	}

}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvRead := csv.NewReader(file)
	for {
		data, err := csvRead.Read()
		if err == io.EOF {
			fmt.Println("*** End of file ***")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		insertQuestions(problemMap, data)
		// fmt.Printf("Question is %v, answer = %v, of the type %T %T\n", data[0], data[1], data[0], data[1])
	}
	fmt.Println(problemMap["1+1"] + problemMap["1+2"])
}
