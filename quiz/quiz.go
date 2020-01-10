package quiz

import (
	"fmt"
	"strings"
)

type problem struct {
	q string
	a string
}

func NewQuiz(records [][]string) {
	var score int
	problems := parseRecords(records)
	for _, p := range problems {
		fmt.Printf("%v = ", p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d\n", score, len(records))
}

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, r := range records {
		ret[i] = problem{
			q: r[0],
			a: strings.TrimSpace(r[1]),
		}
	}

	return ret
}
