package romanToDecimal
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

func RomanToDecimal(roman string) int {
	result :=0

	for _, numeral := range allRomanNumeral {
		for strings.HasPrefix(roman,numeral.Roman) {
			result += numeral.Decimal
			roman = strings.TrimPrefix(roman,numeral.Roman)
		}
	}
	return result
}
