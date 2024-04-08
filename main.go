package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите Первое число Операцию и Второе число : ")
		inputData, _ := reader.ReadString('\n')
		str := strings.ReplaceAll(inputData, " ", "")
		str = strings.ToUpper(str)
		str = strings.TrimSpace(str)
		exam(str)
	}
}

var a, b *int
var math = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"*": func() int { return *a * *b },
	"/": func() int { return *a / *b },
}

var sRoman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

func exam(str string) {
	var mathOperator string
	var cleanStr []string
	for index := range math {
		for _, value := range str {
			if index == string(value) {
				mathOperator += index
				cleanStr = strings.Split(str, mathOperator)
			}
		}
	}

	var foundLetter int
	sNumbers := make([]int, 0)
	sLetters := make([]string, 0)
	if len(mathOperator) != 1 {
		panic("Ошибка")
	} else {
		for _, sim := range cleanStr {
			value, err := strconv.Atoi(sim)
			if err != nil {
				sLetters = append(sLetters, sim)
				foundLetter++
			} else {
				sNumbers = append(sNumbers, value)
			}

		}
	}
	switch foundLetter {
	case 0:
		errorCheck := sNumbers[0] > 0 && sNumbers[0] < 11 && sNumbers[1] > 0 && sNumbers[1] < 11
		if value, ok := math[mathOperator]; ok && errorCheck == true {
			a = &sNumbers[0]
			b = &sNumbers[1]
			fmt.Println(value())

		} else {
			panic("Ошибка")
		}
	case 1:
		panic("Ошибка")
	case 2:
		sLetToNum := make([]int, 0)
		for _, sim := range sLetters {
			if value, ok := sRoman[sim]; ok && value > 0 && value < 11 {
				sLetToNum = append(sLetToNum, value)
			} else {
				panic("Ошибка")
			}
		}
		if value, ok := math[mathOperator]; ok {
			a = &sLetToNum[0]
			b = &sLetToNum[1]
			if value() >= 1 {
				numToLetter(value())
			} else {
				panic("Ошибка")
			}
		}
	}
}

func numToLetter(solution int) {
	numbers := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var str string
	for solution > 0 {
		for i := 0; i < 9; i++ {
			if numbers[i] <= solution {
				str += letters[i]
				solution -= numbers[i]
				break
			}
		}
	}
	fmt.Println(str)
}
