package main

import (
	"fmt"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func printMenue() {
	fmt.Println(string(colorBlue), `
**************************************************
******* Willkommen bei WER WIRD MILLIONÄR? *******
******* Das Quiz besteht aus 10 Fragen ***********
******* Pro Quiz können 2 Joker gewählt werden ***
**************************************************
********** Übersicht aller Commands **************
s       : Quiz starten
a,b,c,d : Antwort auswählen
5       : 50/50 Joker auswählen (2 falsche Antworten werden ausgeblendet)
r       : Retry Joker auswählen (Auf Nummer sicher gehen, bei falscher Antwort bekommt man noch eine Chance)
q       : Quiz beenden`)
	fmt.Println(string(colorReset))
}

func printStartQuiz() {
	fmt.Println(string(colorYellow), "Starte Quiz mit 's'")
	fmt.Println(string(colorReset))
}

func printQuizStarted() {
	fmt.Println(string(colorBlue), "Quiz wurde gestartet, viel Glück!")
	fmt.Println(string(colorReset))
}

func printAskQuestionTitle(number int) {
	fmt.Println(string(colorBlue), "Frage Nr. ", number)
	fmt.Println(string(colorReset))
}

func printAskQuestion(question Question) {
	fmt.Println(string(colorYellow), question.Question)
	fmt.Println(string(colorBlue))

	for i := 0; i < len(question.Answers); i++ {
		answer := question.Answers[i]
		switch i {
		case 0:
			fmt.Println("A: ", answer.Text)
		case 1:
			fmt.Println("B: ", answer.Text)
		case 2:
			fmt.Println("C: ", answer.Text)
		case 3:
			fmt.Println("D: ", answer.Text)
		}
	}

	fmt.Println(string(colorReset))
}

func printCorrectAnswer() {
	fmt.Println(string(colorGreen), "Die Anwort war richtig!")
	if currentLevel < 9 {
		fmt.Println(string(colorGreen), "Nächste Frage..")
	}
	fmt.Println(string(colorReset))
}

func printWrongAnswer() {
	fmt.Println(string(colorRed), "Die Anwort war leider falsch.")
	fmt.Println(string(colorReset))
}

func printGoodLuckNextTime() {
	fmt.Println(string(colorYellow), "Viel Glück beim nächsten Mal!")
	fmt.Println(string(colorReset))
}

func print50ChanceJokerUsed() {
	fmt.Println(string(colorGreen), "Der 50 / 50 Joker wurde eingesetzt!")
	fmt.Println(string(colorReset))
}

func printRetryJokerUsed() {
	fmt.Println(string(colorGreen), "Der Retry Joker wurde eingesetzt!")
	fmt.Println(string(colorReset))
}

func printNewChance() {
	fmt.Println(string(colorGreen), "Neuer Versuch..")
	fmt.Println(string(colorReset))
}

func printSuccess() {
	fmt.Println(string(colorGreen), "Gratulation! Alle 10 Fragen wurden richtig beantwortet.")
	fmt.Println(" Jetzt steht der Million nichts mehr im Wege!")
}

func printExitQuiz() {
	fmt.Println(string(colorYellow), "Quiz wurde beendet!")
	fmt.Println(string(colorReset))
}
