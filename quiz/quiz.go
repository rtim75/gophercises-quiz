package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func StartQuiz(problems [][]string) {
	var score int
	for _, record := range problems {
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
	fmt.Printf("You scored %d out of %d\n", score, len(problems))
}
