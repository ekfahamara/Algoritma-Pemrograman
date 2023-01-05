package main

import "fmt"

var (
	pita  string
	CC    byte
	EOP   bool
	index int
)

func START() {
	index = 0
	CC = pita[index]
	EOP = CC == '.'
}

func ADV() {
	index = index + 1
	CC = pita[index]
	EOP = CC == '.'

}

func DIGIT() bool {
	if CC >= 48 && CC <= 57 {
		return true
	}
	return false
}

func STRIP() {
	for CC == '-' && !EOP {
		ADV()
	}
}

func main() {
	var PRODI string
	var ANS bool

	fmt.Scan(&pita)

	START()

	PRODI = ""
	for CC != '-' && !EOP {
		PRODI += string(CC)
		ADV()
	}
	ANS = false

	if PRODI == "IF" || PRODI == "IT" || PRODI == "SE" {
		ANS = true
		ADV()
		for !EOP && ANS {
			if CC != '.' && DIGIT() {
				ANS = true
				ADV()
			} else {
				STRIP()
				if EOP {
					ANS = false
				}
			}
		}
	} else {
		ANS = false
	}

	fmt.Println(ANS)

}
