package price

// CalculateTotalPrice calculate the total price of a hotel stay.
func CalculateTotalPrice(nights, rate, cityTax int) int {
	return nights*rate + cityTax
}
