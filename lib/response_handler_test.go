package lib

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input    string
	expected FactorDBResponse
}{
	{
		`{"id": "2", "status": "P", "factors": [["2", 1]]}`,
		FactorDBResponse{
			Id:     "2",
			Status: "P",
			Factors: []Factor{
				Factor{Number: 2, Power: 1},
			},
		},
	},
	{
		`{"id": "-1", "status": "Unit", "factors": []}`,
		FactorDBResponse{
			Id:      "-1",
			Status:  "Unit",
			Factors: []Factor{},
		},
	},
	{
		`{"id": "90", "status": "FF", "factors": [["2", 1], ["3", 2], ["5", 1]]}`,
		FactorDBResponse{
			Id:     "90",
			Status: "FF",
			Factors: []Factor{
				Factor{Number: 2, Power: 1},
				Factor{Number: 3, Power: 2},
				Factor{Number: 5, Power: 1},
			},
		},
	},
	{
		`{"id": }`,
		FactorDBResponse{},
	},
}

func TestConvert(t *testing.T) {
	for _, c := range cases {
		got, _ := ConvertToFactorDB([]byte(c.input))
		check, message := CheckEquality(got, c.expected)
		if check == false {
			t.Errorf(message)
		}
	}
}

func CheckEquality(got, expected FactorDBResponse) (bool, string) {
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
