package main

import (
	"./calculator"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	amount, tip := askQuestions()
	calculated := calculator.Calculate(argsToParsedArgs(amount, tip))
	fmt.Printf("%s\n", calculator.Display(calculated))
}

func askQuestions() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the bill amount? ")
	amount, _ := reader.ReadString('\n')
	fmt.Print("What is the tip rate? ")
	tip, _ := reader.ReadString('\n')

	return amount, tip
}

func argsToParsedArgs(amount string, tip string) calculator.ParsedArguments {
	new_amount, new_tip, errors := conversion(amount, tip)

	if len(errors) > 0 {
		for _, v := range errors {
			fmt.Fprintf(os.Stderr, "%v\n", v)
		}
		main()
		// need to exit currnt thread somehow
	}

	return calculator.ParsedArguments{
		BillAmount:    new_amount,
		TipPercentage: new_tip,
	}
}

func conversion(amount string, tip string) (int, int, []error) {
	var errors_slice []error
	new_amount, err := stringToInt(amount)
	if err != nil || new_amount < 0 {
		errors_slice = append(errors_slice, errors.New("There was an error 1!"))
	}
	new_tip, err := stringToInt(tip)
	if err != nil || new_tip < 0 {
		errors_slice = append(errors_slice, errors.New("There was an error 2!"))
	}
	return new_amount, new_tip, errors_slice
}

func stringToInt(str string) (int, error) {
	i, err := strconv.Atoi(strings.Replace(strings.Trim(str, " "), "\n", "", -1))
	return i, err
}
