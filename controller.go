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

	"github.com/thoas/go-funk"
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
		if !isQuizMode || retryJokerUsed > 0 {
			break
		}
		useRetryJoker()
	case "5":
		if !isQuizMode || fiftyChangeJokerUsed {
			break
		}
		use50ChanceJoker50()
	case "q":
		clearTerminal()
		ShutDown()
	}
}

func askQuestion() {
	currentQuestion := questionCatalog[currentLevel]

	printAskQuestionTitle(currentLevel + 1)
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
		currentLevel++
		if currentLevel < 10 {
			askQuestion()
		} else {
			printSuccess()
			ShutDown()
		}
	} else {
		printWrongAnswer()
		if retryJokerUsed == 1 {
			printNewChance()
			retryJokerUsed++
			return
		}
		printGoodLuckNextTime()
		ShutDown()
	}
}

func use50ChanceJoker50() {
	currentQuestion := questionCatalog[currentLevel]
	wrongAnswers := funk.Filter(currentQuestion.Answers, func(a Answer) bool {
		return !a.Correct
	}).([]Answer)

	randomIndex := funk.RandomInt(0, 2)
	for i := 0; i < len(wrongAnswers); i++ {
		if i != randomIndex {
			wrongAnswers[i].Text = ""
		}
	}

	for i, answer := range currentQuestion.Answers {
		if !answer.Correct && !funk.Contains(wrongAnswers, answer) {
			currentQuestion.Answers[i].Text = ""
		}
	}
	print50ChanceJokerUsed()
	fiftyChangeJokerUsed = true
	askQuestion()
}

func useRetryJoker() {
	printRetryJokerUsed()
	retryJokerUsed++
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
