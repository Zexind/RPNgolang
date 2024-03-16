package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

type solver struct{}

func (m *solver) Return() string {
	return "123"
}

type Output struct {
	Result  float64
	Reverse []string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "index.html")
		tmpl := template.Must(template.ParseFiles("index.html"))
		name := r.FormValue("nums")
		result1, _ := evaluateRPN(kalDestroyer3000V2UnichoshitelGovnaPlusTerminatorFekaliy(convertToRPN(getTokens(kalDestroyer3000(name)))))
		result := Output{Result: result1, Reverse: kalDestroyer3000V2UnichoshitelGovnaPlusTerminatorFekaliy(convertToRPN(getTokens(kalDestroyer3000(name))))}
		tmpl.Execute(w, result)
	})

	fmt.Println("YADEBIL")
	l1 := getLinesFromFile("input.txt")
	file, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_TRUNC, 0600)
	defer file.Close()
	for _, text := range l1 {

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
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8081", nil)
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
