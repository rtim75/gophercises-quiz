package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strings"
	"github.com/rtim75/gophercises-quiz/quiz"
)

var problemPath string

func init() {
	flag.StringVar(&problemPath, "f", "problems.csv", "A path to the csv file with problems.")
	flag.Parse()
}

func main() {
	file, err := os.Open(problemPath)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}
	defer file.Close()

	data := make([]byte, 1024)
	length, err := file.Read(data)
	if err != nil {
		log.Fatal("Error reading the file: ", err)
	}

	r := csv.NewReader(strings.NewReader(string(data[:length])))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Error parsing the csv file: ", err)
	}
	quiz.StartQuiz(records)
}
