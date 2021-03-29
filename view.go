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
	colorWhite  = "\033[37m"
)

func printMenue() {
	fmt.Println(string(colorBlue), `
**************************************************
******* Willkommen bei WER WIRD MILLIONÄR? *******
******* Das Quiz besteht aus 10 Fragen ***********
******* Pro Quiz können 2 Joker gewählt werden ***
**************************************************
**** Übersicht Commands ****
s  : Quiz starten
a-d: Antwort auswählen
n  : Nächste Frage
5  : 50/50 Joker auswählen
r  : Retry Joker auswählen
q  : Quiz beenden`)
	fmt.Println(string(colorReset))

	fmt.Println(string(colorRed), "test")
	fmt.Println(string(colorReset))
}

func printStartQuiz() {
	fmt.Println(string(colorYellow), "Starte Quiz mit 's'")
	fmt.Println(string(colorReset))
}

func printExitQuiz() {
	fmt.Println(string(colorYellow), "Quiz wurde beendet!")
	fmt.Println(string(colorReset))
}
