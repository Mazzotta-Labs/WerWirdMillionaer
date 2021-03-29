package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
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

	//for i := 0; i < len(data.Questions); i++ {
	// fmt.Println("Frage: ", i+1)
	// fmt.Println("ID: ", data.Questions[i].Id)
	// fmt.Println("Text:", data.Questions[i].Question)
	// fmt.Println("Difficulty:", data.Questions[i].Difficulty)

	//	for y := 0; y < len(data.Questions[i].Answers); y++ {
	//answer := data.Questions[i].Answers[y]

	// fmt.Println("Antwort: ", y+1)
	// fmt.Println("Text: ", answer.Text)
	// fmt.Println("Korrekt:", answer.Correct)
	//	}
	//}
}

func startQuiz() {
	printStartQuiz()
	prepareQuestionForQuiz()

	for {
		executeCommand()
	}

	// command := askForCommand()
	// if command != "s" {
	// 	os.Exit(3)
	// }
	// fmt.Println("Frage Nummer 01")
}

func prepareQuestionForQuiz() {

	allQuestions := readJsonFile()

	// for i := 0; i < len(allQuestions.Questions); i++ {
	// 	test := allQuestions.Questions[i]
	// 	fmt.Println(test.Id)
	// }

	shuffledQuestion := shuffleQuestions(allQuestions.Questions)

	// for i := 0; i < len(shuffledQuestion); i++ {
	// 	test2 := shuffledQuestion[i]
	// 	fmt.Println(test2.Id)
	// }

	// rand.Shuffle(allQuestions.Questions)
	// easyQuestions := make([]Question, 0, 3)

	// fmt.Println("Easy: ", len(easyQuestions))
	// fmt.Println("Medium: ", len(mediumQuestions))
	// fmt.Println("Hard: ", len(hardQuestions))

	for i := 0; i < len(shuffledQuestion); i++ {
		AddQuestionByDifficulty(shuffledQuestion[i])
	}

	fmt.Println("Easy: ", len(easyQuestions))
	fmt.Println("Medium: ", len(mediumQuestions))
	fmt.Println("Hard: ", len(hardQuestions))

	// }

	//easyQuestions = funk.Contains(allQuestions.Questions, "1")
}

func shuffleQuestions(questions []Question) []Question {
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
	case input == "":
		break
	case input == "q":
		clearTerminal()
		printExitQuiz()
		os.Exit(0)
	}
}

func askQuestion() {

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
