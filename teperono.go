package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("YADEBIL")
	l1 := getLinesFromFile("input.txt")
	for _, text := range l1 {
		file, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0600)
		defer file.Close()
		fmt.Println(kalDestroyer3000(text))

		result, err := evaluateRPN(kalDestroyer3000V2UnichoshitelGovnaPlusTerminatorFekaliy(convertToRPN(getTokens(kalDestroyer3000(text)))))

		fmt.Println(kalDestroyer3000V2UnichoshitelGovnaPlusTerminatorFekaliy(convertToRPN(getTokens(kalDestroyer3000(text)))))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Result:", result)
		result1 := strconv.FormatFloat(result, 'f', -1, 64)
		file.WriteString(result1)
		file.WriteString("\n")
	}
}

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
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid RPN expression")
	}

	return stack[0], nil
}

func getLinesFromFile(adress string) []string {
	fmt.Println("### Read as reader ###")
	f, err := os.Open(adress)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var l []string

	// Чтение файла с ридером
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l = append(l, sc.Text())
	}
	return l
}

func kalDestroyer3000(a string) string {
	var newLine string
	var signesLine = "+-*/^()"
	for _, ch := range a {
		if unicode.IsDigit(ch) || strings.Contains(signesLine, string(ch)) {
			newLine += string(ch)
		}
	}
	return newLine
}

func kalDestroyer3000V2UnichoshitelGovnaPlusTerminatorFekaliy(a []string) []string {
	var signesLine = "()"
	var new []string

	for i := 0; i < len(a); i++ {
		if strings.Contains(signesLine, a[i]) {

		} else {
			new = append(new, a[i])

		}
	}
	return new
}

func getTokens(a string) []string {

	var tokens []string
	re, _ := regexp.Compile(`\d+|[+\/\*\-\^]|[()]`)
	tokens = re.FindAllString(a, -1)
	fmt.Println(tokens)

	return tokens
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
	signs := NewStack[string]()
	var result []string
	var i = 0
	var priority int
	var prevPriority = 0
	prevPrevPriority := 0
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
			result = append(result, signs.Pop())
			prevPrevPriority = getPriority(signs.Peek())
			fmt.Println(prevPrevPriority)
			if prevPrevPriority >= 4 {
				signs.Pop()
			}
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
