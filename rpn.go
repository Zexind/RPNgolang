package main

// 123123123123

import (
	"fmt"
	"math"
	"strconv"
)

func evaluateRPN(tokens []string) (float64, error) {

	stack := make([]float64, 0)

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 32); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid RPN expression")
			}

			// Pop the last two operands
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Perform the operation based on the operator
			switch token {
			case "+":
				stack = append(stack, op1+op2)
			case "-":
				stack = append(stack, op1-op2)
			case "*":
				stack = append(stack, op1*op2)
			case "/":
				if op2 == 0 {
					return 0, fmt.Errorf("division by zero")
				}

				stack = append(stack, op1/op2)
			case "^":
				stack = append(stack, math.Pow(op1, op2))

			default:
				return 0, fmt.Errorf("unknown operator: %s", token)
			}
			fmt.Println(stack)
		}

	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid RPN expression")
	}

	return stack[0], nil
}

func getPriority(token string) int {

	m := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
		")": 4,
		"(": 5,
	}

	n := m[token]
	return n
}

func convertToRPN(tokens []string) []string {
	var signs = NewStack[string]()
	var result []string
	var i = 0
	var priority int
	var prevPriority = 0
	var prevPrevPriority = 0
	for i < len(tokens) {
		priority = getPriority(tokens[i])
		switch priority {
		case 0:
			result = append(result, tokens[i])
		case 1:
			if prevPriority == 2 {
				if signs.Size() == 1 || signs.Size() > 2 {
					result = append(result, signs.Pop())
				} else if signs.Size() == 2 {
					result = append(result, signs.Pop())
					result = append(result, signs.Pop())
				}
			}
			prevPrevPriority = getPriority(signs.Peek())
			fmt.Println(prevPrevPriority)
			if prevPrevPriority == 1 {
				result = append(result, signs.Pop())
			}
			if prevPriority == 3 {
				result = append(result, signs.Pop())
			}
			signs.Push(tokens[i])
		case 2:
			prevPrevPriority = getPriority(signs.Peek())
			fmt.Println(prevPrevPriority)
			if prevPrevPriority == 1 {
				signs.Push(tokens[i])
				break
			} else if prevPriority >= 2 {
				result = append(result, signs.Pop())
			}
			signs.Push(tokens[i])
		case 3:
			signs.Push(tokens[i])
		case 4:
			prevPrevPriority = getPriority(signs.Peek())
			fmt.Println(prevPrevPriority)
			for signs.Peek() != "(" && !signs.IsEmpty() {
				result = append(result, signs.Pop())
			}
			signs.Pop()
		case 5:
			signs.Push(tokens[i])
		}
		if priority != 0 {
			prevPriority = priority
		}
		i++
		fmt.Println(signs)
		fmt.Println(result)
	}
	for signs.Size() > 0 {
		result = append(result, signs.Pop())
	}
	fmt.Println(result)
	return result
}
