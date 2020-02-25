package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"

	"github.com/rtim75/gophercises-quiz/quiz"
)

var problemPath string

var timer int

func init() {
	flag.StringVar(&problemPath, "f", "problems.csv", "A path to the csv file with problems.")
	flag.IntVar(&timer, "t", 30, "A total time you have to solve all problems.")
	flag.Parse()
}

func main() {
	file, err := os.Open(problemPath)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error parsing the csv file: ", err)
	}
	quiz.NewQuiz(records, timer)
}
