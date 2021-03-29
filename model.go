package main

import "os"

var easyQuestions []Question
var mediumQuestions []Question
var hardQuestions []Question
var questionCatalog []Question
var currentLevel int = 0

type Questions struct {
	Questions []Question
}

type Question struct {
	Id         string   `json:"id"`
	Question   string   `json:"question"`
	Difficulty int      `json:"difficulty"`
	Answers    []Answer `json:"answers"`
}

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

func AddQuestionByDifficulty(question Question) {
	switch question.Difficulty {
	case 1:
		if len(easyQuestions) < 3 {
			easyQuestions = append(easyQuestions, question)
		}
	case 2:
		if len(mediumQuestions) < 4 {
			mediumQuestions = append(mediumQuestions, question)
		}
	case 3:
		if len(hardQuestions) < 3 {
			hardQuestions = append(hardQuestions, question)
		}
	}
}

func SetupQuestionsForQuiz() {
	questionCatalog = append(questionCatalog, easyQuestions...)
	questionCatalog = append(questionCatalog, mediumQuestions...)
	questionCatalog = append(questionCatalog, hardQuestions...)
}

func ShutDown() {
	printExitQuiz()
	os.Exit(0)
}
