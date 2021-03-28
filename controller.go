package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
		fmt.Println("Frage: ", i+1)
		fmt.Println("ID: ", data.Questions[i].Id)
		fmt.Println("Text:", data.Questions[i].Question)
		fmt.Println("Difficulty:", data.Questions[i].Difficulty)
		fmt.Println("\n")

		for y := 0; y < len(data.Questions[i].Answers); y++ {
			answer := data.Questions[i].Answers[y]

			fmt.Println("Antwort: ", y+1)
			fmt.Println("Text: ", answer.Text)
			fmt.Println("Korrekt:", answer.Correct)
			fmt.Println("\n")
		}
	}
}
