package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func clearTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func readJsonFile() {
	file, _ := ioutil.ReadFile("db/QuestionCatalog.json")
	data := Questions{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Questions); i++ {
		// fmt.Println("Frage: ", i+1)
		// fmt.Println("ID: ", data.Questions[i].Id)
		// fmt.Println("Text:", data.Questions[i].Question)
		// fmt.Println("Difficulty:", data.Questions[i].Difficulty)

		for y := 0; y < len(data.Questions[i].Answers); y++ {
			//answer := data.Questions[i].Answers[y]

			// fmt.Println("Antwort: ", y+1)
			// fmt.Println("Text: ", answer.Text)
			// fmt.Println("Korrekt:", answer.Correct)
		}
	}
}

func startQuiz() {
	printStartQuiz()

	for {
		executeCommand()
	}

	// command := askForCommand()
	// if command != "s" {
	// 	os.Exit(3)
	// }
	// fmt.Println("Frage Nummer 01")
}

func executeCommand() {
	command := askForCommand()
	parseCommand(command)
}

func parseCommand(input string) {
	switch {
	case input == "s":
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

func askForCommand() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSpace(response)
	return response
}
