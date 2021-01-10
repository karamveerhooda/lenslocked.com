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
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in 'Question, answer' format")
	flag.Parse()
	_ = csvFilename

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

	//Now we need to iterate through all the problem and get the user input and check if user correct
	//
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// This function is concerting a 2D slice in to a struct of Q and A format
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
