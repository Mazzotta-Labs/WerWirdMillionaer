package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var isQuizMode = false

func clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func readJsonFile() Questions {
	file, _ := ioutil.ReadFile("db/QuestionCatalog.json")
	data := Questions{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}

func startQuiz() {
	printStartQuiz()
	prepareQuestionsForQuiz()

	for {
		executeCommand()
	}
}

func prepareQuestionsForQuiz() {
	allQuestions := readJsonFile()
	shuffledQuestions := shuffleQuestions(allQuestions.Questions)

	for i := 0; i < len(shuffledQuestions); i++ {
		AddQuestionByDifficulty(shuffledQuestions[i])
	}

	SetupQuestionsForQuiz()
}

func shuffleQuestions(questions []Question) []Question {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	return questions
}

func executeCommand() {
	command := askForCommand()
	parseCommand(command)
}

func parseCommand(input string) {
	switch strings.ToLower(input) {
	case "s":
		if isQuizMode {
			break
		}

		isQuizMode = true
		printQuizStarted()
		askQuestion()
	case "a", "b", "c", "d":
		if !isQuizMode {
			break
		}

		checkAnswer(input)
	case "r":
		break
	case "5":
		break
	case "q":
		clearTerminal()
		ShutDown()
	}
}

func askQuestion() {
	currentLevel++

	printAskQuestionTitle(currentLevel)
	currentQuestion := questionCatalog[currentLevel]

	printAskQuestion(currentQuestion)
}

func checkAnswer(input string) {
	currentQuestion := questionCatalog[currentLevel]
	var index int

	switch strings.ToLower(input) {
	case "a":
		index = 0
	case "b":
		index = 1
	case "c":
		index = 2
	case "d":
		index = 3
	}

	if currentQuestion.Answers[index].Correct {
		printCorrectAnswer()
		askQuestion()
	} else {
		printWrongAnswer()
		ShutDown()
	}
}

func askForCommand() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSpace(response)
	return response
}
