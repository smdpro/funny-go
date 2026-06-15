package decimalToRoman

import (
	"fmt"
	"testing"
)

func TestDecimalToRoman(t *testing.T) {

	cases := []RomanNumeral{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{900, "CM"},
		{1000, "M"},
		{1994, "MCMXCIV"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d should be converted to %q", c.Decimal, c.Roman), func(t *testing.T) {
			result := DecimalToRoman(c.Decimal)
			if result != c.Roman {
				t.Errorf(`result: "%s",expected: "%s`, result, c.Roman)
			}
		})
	}

}
