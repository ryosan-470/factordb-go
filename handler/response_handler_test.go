package handler

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input    string
	expected FactorDB
}{
	{
		`{"id": "2", "status": "P", "factors": [["2", 1]]}`,
		FactorDB{Id: "2", Status: "P", Factors: []Factor{Factor{Number: 2, Power: 1}}},
	},
}

func TestConvert(t *testing.T) {
	for _, c := range cases {
		got := ConvertToFactorDB([]byte(c.input))
		check, message := CheckEquality(got, c.expected)
		if check == false {
			t.Errorf(message)
		}
	}
}

func CheckEquality(got, expected FactorDB) (bool, string) {
	if got.Status != expected.Status {
		message := "Both status must be equality"
		return false, message
	}

	if got.Id != expected.Id {
		message := "Both id must be equality"
		return false, message
	}

	if len(got.Factors) != len(expected.Factors) {
		message := "Both factors must be same size"
		return false, message
	}

	for i, expect := range expected.Factors {
		g := got.Factors[i]
		if g.Number != expect.Number || g.Power != expect.Power {
			message := `The element of Factors must be same\n
got.Factors = %v\n
expect = %v\n`
			message = fmt.Sprintf(message, g, expect)
			return false, message
		}
	}
	return true, ""
}
