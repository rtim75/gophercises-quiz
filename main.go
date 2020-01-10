package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	var score int
	for _, record := range records {
		input_reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%v = ", record[0])
		input, err := input_reader.ReadString('\n')
		answer := strings.TrimSuffix(input, "\n")
		if err != nil {
			log.Fatal("Error reading from stdin: ", err)
		}
		if answer == record[1] {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d", score, len(records))
}
