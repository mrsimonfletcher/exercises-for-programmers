package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := input()
	newResponse := create_new_response(input)
	output(newResponse)
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	return trimArg(name)
}

func create_new_response(input string) string {
	s := []string{"Hello", input, "nice to meet you!"}
	return strings.Join(s, ", ")
}

func output(str string) {
	fmt.Printf(str)
}

func trimArg(str string) string {
	return strings.Replace(strings.Trim(str, " "), "\n", "", -1)
}
