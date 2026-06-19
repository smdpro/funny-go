package romanToDecimal

import (
	"fmt"
	"testing"
)

func TestRomanToDecimal(t *testing.T) {

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
		t.Run(fmt.Sprintf("%q should be converted to %d", c.Roman, c.Decimal), func(t *testing.T) {
			result := RomanToDecimal(c.Roman)
			if result != c.Decimal {
				t.Errorf(`result: "%d",expected: "%s`, result, c.Roman)
			}
		})
	}

}
