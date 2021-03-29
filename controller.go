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
	switch {
	case input == "s":
		printQuizStarted()
		askQuestion()
		break
	case input == "a" || input == "b" || input == "c" || input == "d":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "":
		break
	case input == "q":
		clearTerminal()
		printExitQuiz()
		os.Exit(0)
	}
}

func askQuestion() {
	currentLevel++

	printAskQuestionTitle(currentLevel)
	currentQuestion := questionCatalog[currentLevel]

	printAskQuestion(currentQuestion)
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
