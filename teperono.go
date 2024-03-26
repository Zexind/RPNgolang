package main

import (
	"bufio"
	"fmt"

	// "net/http"

	"os"
	"regexp"
	"strconv"
	"strings"

	// "text/template"
	"unicode"
)

type Output struct {
	Result  float64
	Reverse []string
}

func main() {

	/*	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("index.html"))
			name := r.FormValue("nums")
			result1, _ := evaluateRPN(RemoveInArray(convertToRPN(getTokens(RemoveInString(name)))))
			result := Output{Result: result1, Reverse: RemoveInArray(convertToRPN(getTokens(RemoveInString(name))))}
			tmpl.Execute(w, result)
		})
	*/

	fmt.Println("YADEBIL")
	l1 := getLinesFromFile("input.txt")
	file, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_TRUNC, 0600)
	defer file.Close()
	for _, text := range l1 {
		result, err := evaluateRPN(RemoveInArray(convertToRPN(getTokens(RemoveInString(text)))))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Result:", result)
		result1 := strconv.FormatFloat(result, 'f', -1, 64)
		file.WriteString(result1)
		file.WriteString("\n")
	}

	// fmt.Println("Server is listening...")
	// http.ListenAndServe(":8081", nil)
}

func getLinesFromFile(adress string) []string {
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

func RemoveInString(a string) string {
	var newLine string
	var signesLine = "+-*/^()"
	for _, ch := range a {
		if unicode.IsDigit(ch) || strings.Contains(signesLine, string(ch)) {
			newLine += string(ch)
		}
	}
	return newLine
}

func RemoveInArray(a []string) []string {
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
	slice := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		slice = append(slice, tokens[i+1:]...)
		if i != len(tokens)-1 && tokens[i] == ")" && tokens[i+1] == "(" {
			tokens = append(tokens[:i+1], "*")
			tokens = append(tokens, slice...)
		}
		if i != len(tokens)-1 && getPriority(tokens[i]) == 0 && tokens[i+1] == "(" {
			tokens = append(tokens[:i+1], "*")
			tokens = append(tokens, slice...)
		}
		if i != len(tokens)-1 && getPriority(tokens[i+1]) == 0 && tokens[i] == ")" {
			tokens = append(tokens[:i+1], "*")
			tokens = append(tokens, slice...)
		}
		slice = slice[:0]
	}
	return tokens
}
