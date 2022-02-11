package roman_test

import (
	"fmt"
	"testing"

	"git.ivxs.uk/daniel.perez/ca-golang-workshop/app/pkg/roman"
)

func TestRomanNumerals(t *testing.T) {
	suite := []struct {
		Roman   string
		Decimal int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"VI", 6},
		{"VIII", 8},
		{"IX", 9},
		{"XLIX", 49},
		{"CXX", 120},
		{"XXXIX", 39},
		{"MMVIII", 2008},
		{"MMXXII", 2022},
	}

	for _, test := range suite {
		t.Run(fmt.Sprintf(`"%s" equals %d`, test.Roman, test.Decimal), func(t *testing.T) {
			converted := roman.ToDecimal(test.Roman)
			if test.Decimal != converted {
				t.Errorf(`Expected %v to equal %d`, converted, test.Decimal)
			}
		})
	}
}
