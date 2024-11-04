// Реализовать функцию func Calc(expression string) (float64, error)
// expression - строка-выражение состоящее из односимвольных идентификаторов и знаков арифметических действий
// Входящие данные - цифры(рациональные), операции +, -, *, /, операции приоритезации ( и ) В случае ошибки записи выражения функция выдает ошибку.

// Сохраните этот код себе на github. Он понадобится вам при выполнении финальных заданий следующих модулей

package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func Operator(char byte) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func Priority(op byte) int {
	if op == '*' || op == '/' {
		return 2
	}
	if op == '+' || op == '-' {
		return 1
	}
	return 0
}

func ApplyOperator(a, b float64, operator byte) (float64, error) {
	switch operator {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, errors.New("zero")
		}
		return a / b, nil
	}
	return 0, errors.New("inv operator")
}

func Calc(expression string) (float64, error) {
	var numStack []float64
	var opStack []byte

	i := 0
	for i < len(expression) {
		char := expression[i]

		// Skip whitespace
		if unicode.IsSpace(rune(char)) {
			i++
			continue
		}

		if unicode.IsDigit(rune(char)) || char == '.' {
			numStart := i
			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				i++
			}
			num, err := strconv.ParseFloat(expression[numStart:i], 64)
			if err != nil {
				return 0, fmt.Errorf("inv num: %v", err)
			}
			numStack = append(numStack, num)
			continue
		}

		if Operator(char) {
			for len(opStack) > 0 && Priority(opStack[len(opStack)-1]) >= Priority(char) {
				if len(numStack) < 2 {
					return 0, errors.New("inv expression")
				}
				b := numStack[len(numStack)-1]
				a := numStack[len(numStack)-2]
				numStack = numStack[:len(numStack)-2]
				result, err := ApplyOperator(a, b, opStack[len(opStack)-1])
				if err != nil {
					return 0, err
				}
				opStack = opStack[:len(opStack)-1]
				numStack = append(numStack, result)
			}
			opStack = append(opStack, char)
			i++
			continue
		}

		if char == '(' {
			opStack = append(opStack, char)
		} else if char == ')' {
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				if len(numStack) < 2 {
					return 0, errors.New("invalid expression")
				}
				b := numStack[len(numStack)-1]
				a := numStack[len(numStack)-2]
				numStack = numStack[:len(numStack)-2]
				result, err := ApplyOperator(a, b, opStack[len(opStack)-1])
				if err != nil {
					return 0, err
				}
				opStack = opStack[:len(opStack)-1]
				numStack = append(numStack, result)
			}
			if len(opStack) == 0 || opStack[len(opStack)-1] != '(' {
				return 0, errors.New("mismatched parentheses")
			}
			opStack = opStack[:len(opStack)-1]
		}
		i++
	}

	for len(opStack) > 0 {
		if len(numStack) < 2 {
			return 0, errors.New("inv expression")
		}
		b := numStack[len(numStack)-1]
		a := numStack[len(numStack)-2]
		numStack = numStack[:len(numStack)-2]
		result, err := ApplyOperator(a, b, opStack[len(opStack)-1])
		if err != nil {
			return 0, err
		}
		opStack = opStack[:len(opStack)-1]
		numStack = append(numStack, result)
	}

	if len(numStack) != 1 {
		return 0, errors.New("inv expression")
	}

	return numStack[0], nil
}
