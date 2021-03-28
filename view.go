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
# 1. ADD ANOTHER BOOK IN LIBRARY
# 2. REMOVE A BOOK FROM A LIBRARY
# 3. CHECK AVAILABILITY
# 4. LEND A BOOK
# 5. RETURN A BOOK
# 6. VIEW ALL BOOKS
#
# q. TERMINATE BOOK LIBRARY APP`)
	fmt.Println(string(colorReset))

	fmt.Println(string(colorRed), "test")
	fmt.Println(string(colorReset))
}
