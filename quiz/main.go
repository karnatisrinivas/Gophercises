package main
import (
	"encoding/csv"
	"fmt"
	"flag"
	"os"
	"strings"
	"time"

)
func main(){
	fileName := flag.String("csv", "problems.csv", "file name of the csv file")
	// Flag for filename
	timeSeconds := flag.Int("time", 15, "Time limit for the quiz")
	flag.Parse()


	// fmt.Println(*fileName) Default, just passing filename gives address

	file, err := os.Open(*fileName) 
	if err!= nil {
		fmt.Println("File not found or it can't be opened")
	}
	
	content := csv.NewReader(file)
	lines, err := content.ReadAll()
	if err != nil {
		fmt.Println("Content is the csv file is not readable")
	}
	problems := parseData(lines)

	timer := time.NewTimer(time.Duration(*timeSeconds)*time.Second)

	correctCount := 0
	for _, p := range problems{
		fmt.Printf("%s = \n", p.q)
		answerChannel :=make(chan string)
		go func(){
	
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You got %v correct out of %v\n", correctCount, len(lines))
			return
		case answer := <-answerChannel:
			if answer==p.a{
				correctCount++
			}
		}
	}
	fmt.Printf("\nYou got %v correct out of %v\n", correctCount, len(lines))
}

func parseData(lines [][]string) []problem {
	ret:= make([]problem, len(lines))
	for i, line := range lines{
		ret[i]= problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}