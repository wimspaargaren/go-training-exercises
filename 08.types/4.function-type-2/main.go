package main

type CalcType int

const (
	// CalcTypeSum Add a type for Subtract and Multiplication.
	CalcTypeSum CalcType = iota // iota: https://go.dev/wiki/Iota
)

type Calc func(x, y int) int

func main() {
	// Call the GetCalculationFunction with the different CalcTypes and execute them with the following variables
	_, _ = 42, 42
	// Execute the test to verify your solution
}

func GetCalculationFunction(calcType CalcType) Calc {
	// Implement the GetCalculationFunction
	// Hint: Use a switch statement to return the correct function
	return nil
}
