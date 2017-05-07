package calculator

import "testing"

func TestDisplay(t *testing.T) {
	actual := Display(display{
		tip:   0.00,
		total: 10.00,
	})
	expected := "Tip: $0.00\nTotal: $10.00\n"

	if actual != expected {
		t.Errorf("Display test expected [%s], actual [%s]", expected, actual)
	}
}

func TestCalculate(t *testing.T) {
	actual := Calculate(ParsedArguments{
		billAmount:    200,
		tipPercentage: 15,
	})

	expected := display{
		tip:   30.00,
		total: 230.00,
	}

	if actual != expected {
		t.Errorf("Display test expected [%#v], actual [%#v]", expected, actual)
	}
}
