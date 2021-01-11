//Q:
package main

//Notes: 1. In comma separated csv file , all the lines should have same format like
// 1+1,2
// 2+2,4
// 3+2,5,10 will result in error
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in 'Question, answer' format")
	timelimit := flag.Int("limit", 30, "the time limit of quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error in opening the csv file: %s\n", *csvFilename))
	}
	//create a csv reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the csv file.")
	}
	problems := parseLines(lines)
	//timer comes after parselines because all set up needs to be completed before timer starts
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	//Now we need to iterate through all the problem and get the user input and check if user correct
	//
	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// This function is converting a 2D slice in to a struct of Q and A format.  Or more readable format
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

//Problem func will allow us to choose any kind of file like JSON or html. So declaring a struct of format question and answer
type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // our application have error
}
