package quiz

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func NewQuiz(records [][]string, t int, shuffle bool) {
	var score int
	problems := parseRecords(records, shuffle)
	timer := time.NewTimer(time.Duration(t) * time.Second)
	for _, p := range problems {
		fmt.Printf("%v = ", p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
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

func parseRecords(records [][]string, shuffle bool) []problem {
	ret := make([]problem, len(records))
	for i, r := range records {
		ret[i] = problem{
			q: r[0],
			a: strings.TrimSpace(r[1]),
		}
	}

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(ret), func(i, j int) { ret[i], ret[j] = ret[j], ret[i] })
	}

	return ret
}
