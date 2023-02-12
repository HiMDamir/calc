package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arabicToRoman = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}
var romanToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func arabicToRomanNumeral(arabic int) string {
	if arabic > 10 || arabic < 1 {
		fmt.Println("Ошибка: Ввод должен быть числом от 1 до 10")
		return ""
	}
	return arabicToRoman[arabic]
}
func romanToArabicNumeral(roman string) int {
	arabic, exists := romanToArabic[roman]
	if !exists {
		fmt.Println("Ошибка: Неверная римская цифра")
	}
	return arabic
}
func isArabic(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
func isRoman(input string) bool {
	_, exists := romanToArabic[input]
	return exists
}
func checkInput(input string) (int, bool) {
	if isArabic(input) {
		arabic, _ := strconv.Atoi(input)
		return arabic, true
	} else if isRoman(input) {
		return romanToArabicNumeral(input), false
	}
	fmt.Println("Ошибка: Ввод должен быть либо арабскими, либо римскими цифрами")
	//os.Exit(1)
	return 0, false
}
func evaluateExpression(firstOperand int, secondOperand int, operator string) int {
	switch operator {
	case "+":
		return firstOperand + secondOperand
	case "-":
		return firstOperand - secondOperand
	case "*":
		return firstOperand * secondOperand
	case "/":
		return firstOperand / secondOperand
	default:
		fmt.Println("Ошибка ввода данных")
		return 0
	}
}
func main() {
	fmt.Print("Введите значение: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	var output string
	for _, char := range input {
		if char != ' ' {
			output += string(char)
		}
	}
	input = output
	operator := ""
	operationCount := 0
	for _, char := range input {
		if string(char) == "+" || string(char) == "-" || string(char) == "/" || string(char) == "*" {
			operator = string(char)
			operationCount++
		}
	}
	if operationCount > 1 {
		fmt.Println("Возможна только одна операция вычесления")
		return
	} else if operator == "" {
		fmt.Println("Ошибка: Недопустимый оператор")
		return
	}
	split := strings.Split(input, operator)
	firstInput := split[0]
	secondInput := split[1]
	secondInput = strings.TrimSpace(secondInput)
	firstOperand, isFirstOperandArabic := checkInput(firstInput)
	arabicToRomanNumeral(firstOperand)
	secondOperand, isSecondOperandArabic := checkInput(secondInput)
	arabicToRomanNumeral(secondOperand)
	result := evaluateExpression(firstOperand, secondOperand, operator)
	if isFirstOperandArabic && isSecondOperandArabic {
		fmt.Printf("Результат (арабский): %d\n", result)
	} else if !isFirstOperandArabic && !isSecondOperandArabic {
		fmt.Printf("Результат (римский): %s\n", arabicToRomanNumeral(result))
	} else {
		fmt.Println("Ошибка: Оба выражения должны быть либо арабскими, либо римскими цифрами")
	}
}
