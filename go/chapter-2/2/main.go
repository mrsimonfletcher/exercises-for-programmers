package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := input()
	newResponse := createNewResponse(input)
	output(newResponse)
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	return trimArg(name)
}

func createNewResponse(input string) string {
	s := []string{input, "has", strconv.Itoa(len(input)), "characters"}
	return strings.Join(s, " ")
}

func output(str string) {
	fmt.Printf("%s.", str)
}

func trimArg(str string) string {
	return strings.Replace(strings.Trim(str, " "), "\n", "", -1)
}
