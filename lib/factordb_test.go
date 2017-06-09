package lib

import (
	"fmt"
	"testing"
)

var result FactorDBResponse = FactorDBResponse{
	Id:     "1",
	Status: "Unit",
}

var db FactorDB = FactorDB{
	Number: 1,
	Result: result,
}

var error_db FactorDB = FactorDB{
	Number: 1,
}

func TestEmpty(t *testing.T) {
	var cases = []struct {
		input    FactorDB
		expected bool
	}{
		{
			FactorDB{
				Number: 1,
				Result: FactorDBResponse{},
			},
			true,
		},
		{
			db,
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

func TestConnect(t *testing.T) {
	var cases = []struct {
		input    FactorDB
		expected error
	}{
		{
			FactorDB{Number: 10},
			nil,
		},
	}

	for _, c := range cases {
		got := c.input.Connect()
		if got != c.expected {
			t.Errorf("Test Connect: %v", c.input)
		}
	}
}

func TestGetId(t *testing.T) {
	var cases = []struct {
		input    FactorDB
		expected string
	}{
		{
			db,
			"1",
		},
		{
			error_db,
			"",
		},
	}

	for _, c := range cases {
		got, _ := c.input.GetId()
		if got != c.expected {
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
			db,
			"Unit",
		},
		{
			error_db,
			"",
		},
	}

	for _, c := range cases {
		got, _ := c.input.GetStatus()
		if got != c.expected {
			t.Errorf("got = %s expected = %s", got, c.expected)
		}
	}
}

func TestGetFactorList(t *testing.T) {
	var testFactors = []Factor{
		Factor{Number: 2, Power: 1},
		Factor{Number: 3, Power: 2},
		Factor{Number: 5, Power: 1},
	}

	var cases = []struct {
		input    FactorDB
		expected []int
	}{
		{
			FactorDB{
				Number: 1,
				Result: FactorDBResponse{
					Status:  "Unit",
					Factors: testFactors,
				},
			},
			[]int{2, 3, 3, 5},
		},
		{
			FactorDB{
				Number: 1,
				Result: FactorDBResponse{
					Status: "Unit",
					Factors: []Factor{
						Factor{},
					},
				},
			},
			[]int{},
		},
		{
			error_db,
			[]int{},
		},
	}

	for _, c := range cases {
		got, _ := c.input.GetFactorList()
		if !CheckEqualitySlice(got, c.expected) {
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
