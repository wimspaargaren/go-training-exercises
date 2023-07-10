package main

type CalcType int

const (
	// Add a type for Subtract and Multiplication
	CalcTypeSum CalcType = iota
)

type Calc func(x, y int) int

func main() {
	// Call the GetCalculationFunction with the different CalcTypes and execute them with the following variables
	_, _ = 42, 42
	// Execute the test to verify your solution
}
