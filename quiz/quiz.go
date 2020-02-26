package quiz

import (
	"fmt"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func NewQuiz(records [][]string, t int) {
	var score int
	problems := parseRecords(records)
	timer := time.NewTimer(time.Duration(t) * time.Second)
	for _, p := range problems {
		fmt.Printf("%v = ", p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			fmt.Print(&answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYour time is over.\nYou scored %d out of %d\n", score, len(records))
			return
		case answer := <-answerCh:
			if answer == p.a {
				score++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d\n", score, len(records))
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
