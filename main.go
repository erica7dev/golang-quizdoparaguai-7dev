package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"encoding/csv"
)

//1. ler csv
func main() {
	fName := flag.String("f", "main.go")
	timer := flag.Int("t", 30, "timer of quiz")
	flag.Parse()
	problems, err := problemPuller(*fName)

	if err != nil {
		exit(fmt.Sprintf("Failed to parse the CSV file: %s", err.Error()))
	}

	CorrectAns := 0

	tObj := timer.newTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)

	problemLoop:

		for i, p := range problems {
			fmt.Printf("Problem #%d: %s\n", i+1, p.q)
			go func () {
				fmt.Scanf("%s", &answer)
				ansC <- answer
			}()

			select{
			case <-tObj.C:
				fmt.Println("Time's up!")
				break problemLoop
			case iAns := <-ansC:
				if iAns == p.a {
					CorrectAns++
				}
				if i == len(problems)-1{
					close(ansC)
				}
			}
}

fmt.Printf("Your result is %d out of %d\n", CorrectAns, len(problems))
fmt.Printf("Press enter to exit")
<- ansC

func problemPuller(fName string) ([]problem, error) {
	if fObj, err := os.Open(fileName); err == nil {
		csvR := csv.NewReader(fObj)

		if clines, err := csvR.ReadAll(); err == nil {
			return parseProblem(clines), nil
	}else{
		return nil, fmt.Errorf("Failed to parse the CSV file: %s", filename, err.Error
	} 
} else {
	return nil, fmt.Errorf("Error to opening %s file; %s", filename, err.Error())
	}
}

func parseProblem(lines [][] string) []problem {
	r := make([]problem, len(lines))
	for i :=0; i < len(lines); i++ {
		r[i] = problem{
			q: lines[i][0],
			a: lines[i][1],
		}
	}
}

type problem struct {
	q string
	a string
}

func exit (){
	fmt.Println("Exiting...")
	os.Exit(1)
}

/*
f, _ := os.Open("quiz.csv")
	r := csv.NewReader(f)

	for {
		record , err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
*/
