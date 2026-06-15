package decimalToRoman

import "strings"

type RomanNumeral struct {
	Decimal int
	Roman   string
}

var allRomanNumeral = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func DecimalToRoman(num int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumeral {
		for num >= numeral.Decimal {
			result.WriteString(numeral.Roman)
			num -= numeral.Decimal
		}
	}
	return result.String()
}
