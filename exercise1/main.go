package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Prompt string
	Answer string
}

func loadQuestions(fileName string, shuffle bool) []Question {
	var questions []Question

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error loading file")
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error parsing file")
	}

	for _, line := range data {
		var question Question
		for j, field := range line {
			if j == 0 {
				question.Prompt = field
			} else {
				question.Answer = field
			}
		}
		questions = append(questions, question)
	}

	if shuffle {
		currentIndex := len(questions) - 1
		for i := 0; i < currentIndex; currentIndex-- {
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(currentIndex)

			questions[currentIndex], questions[randomIndex] = questions[randomIndex], questions[currentIndex]
		}
	}

	return questions
}

func main() {
	filePtr := flag.String("file", "problems.csv", "CSV file with questions")
	timerPtr := flag.Duration("time", 30*time.Second, "Time limit for quiz")
	shufflePtr := flag.Bool("shuffle", false, "If the questions should be shuffled")
	flag.Parse()

	var questions []Question = loadQuestions(*filePtr, *shufflePtr)

	reader := bufio.NewReader(os.Stdin)

	var points int = 0

	fmt.Println("Lets start the quiz!")
	fmt.Println("Click enter when you are ready")
	reader.ReadString('\n')

	timer := time.NewTimer(*timerPtr)

	go func() {
		<-timer.C
		fmt.Printf("\nTimes up! You got %v/%v", points, len(questions))
		os.Exit(0)
	}()

	for _, question := range questions {
		fmt.Print(question.Prompt, " = ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error parsing input")
		}
		trimmedInput := strings.Replace(input, "\n", "", -1)
		if trimmedInput == question.Answer {
			points += 1
		}
	}

	timer.Stop()
	fmt.Printf("Good job, you got %v/%v", points, len(questions))
	os.Exit(0)
}
