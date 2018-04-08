package calculator

import "fmt"

type ParsedArguments struct {
	BillAmount    int
	TipPercentage int
}

type display struct {
	tip   float32
	total float32
}

func Display(arg display) string {
	return fmt.Sprintf("Tip: $%.2f\nTotal: $%.2f\n", arg.tip, arg.total)
}

func Calculate(parsedArgs ParsedArguments) display {
	tip := (parsedArgs.BillAmount / 100) * parsedArgs.TipPercentage

	return display{
		tip:   float32(tip),
		total: float32(parsedArgs.BillAmount + tip),
	}
}
