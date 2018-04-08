package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type personAndQuote struct {
	question string
	person   string
}

func main() {
	input := input()
	personAndQ := quoteAndPerson(input)
	response := createNewResponse(personAndQ)
	fmt.Printf("%s", response)
}

func quoteAndPerson(str string) personAndQuote {
	questionMarkIndex := strings.Index(str, "?")

	question := str[:questionMarkIndex]
	person := str[questionMarkIndex:]
	return personAndQuote{
		question: question,
		person:   person,
	}
}

func createNewResponse(arg personAndQuote) string {
	s := []string{arg.person, " says, ", "\"", arg.question, "\""}
	return strings.Join(s, "")
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the quote? ")
	name, _ := reader.ReadString('\n')
	return trimArg(name)
}

func trimArg(str string) string {
	return strings.Replace(strings.Trim(str, " "), "\n", "", -1)
}
