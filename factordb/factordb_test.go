package factordb

import (
	"fmt"
	"testing"

	"github.com/ryosan-470/factordb-go/handler"
)

func TestEmpty(t *testing.T) {
	var cases = []struct {
		input    FactorDB
		expected bool
	}{
		{
			FactorDB{
				Number: 1,
				Result: handler.FactorDbResponse{},
			},
			true,
		},
		{
			FactorDB{
				Number: 1,
				Result: handler.FactorDbResponse{
					Id:     "-1",
					Status: "Unit",
				},
			},
			false,
		},
	}

	for _, c := range cases {
		got := c.input.Empty()
		if got != c.expected {
			t.Errorf("Test Empty: %v", c.input)
		}
	}
}

func TestGetId(t *testing.T) {
	var cases = []struct {
		input    FactorDB
		expected string
	}{
		{
			FactorDB{
				Number: 1,
				Result: handler.FactorDbResponse{
					Id:     "1",
					Status: "Unit",
				},
			},
			"1",
		},
	}

	for _, c := range cases {
		got, err := c.input.GetId()
		if got != c.expected || err != nil {
			t.Errorf("got = %s expected = %s", got, c.expected)
		}
	}
}

func TestGetStatus(t *testing.T) {
	var cases = []struct {
		input    FactorDB
		expected string
	}{
		{
			FactorDB{
				Number: 1,
				Result: handler.FactorDbResponse{
					Id:     "1",
					Status: "Unit",
				},
			},
			"Unit",
		},
	}

	for _, c := range cases {
		got, err := c.input.GetStatus()
		if got != c.expected || err != nil {
			t.Errorf("got = %s expected = %s", got, c.expected)
		}
	}
}

func TestGetFactorList(t *testing.T) {
	var testFactors = []handler.Factor{
		handler.Factor{Number: 2, Power: 1},
		handler.Factor{Number: 3, Power: 2},
		handler.Factor{Number: 5, Power: 1},
	}

	var cases = []struct {
		input    FactorDB
		expected []int
	}{
		{
			FactorDB{
				Number: 1,
				Result: handler.FactorDbResponse{
					Status:  "Unit",
					Factors: testFactors,
				},
			},
			[]int{2, 3, 3, 5},
		},
		{
			FactorDB{
				Number: 1,
				Result: handler.FactorDbResponse{
					Status: "Unit",
					Factors: []handler.Factor{
						handler.Factor{},
					},
				},
			},
			[]int{},
		},
	}

	for _, c := range cases {
		got, err := c.input.GetFactorList()
		if !CheckEqualitySlice(got, c.expected) || err != nil {
			t.Errorf("got = %v expected = %v", got, c.expected)
		}
	}
}

func CheckEqualitySlice(got, expected []int) bool {
	if len(got) != len(expected) {
		return false
	}

	for i, g := range got {
		if g != expected[i] {
			fmt.Println(g, expected[i])
			return false
		}
	}
	return true
}
