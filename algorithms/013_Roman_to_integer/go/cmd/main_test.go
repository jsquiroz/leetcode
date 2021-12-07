package main

import "testing"

type values struct {
	letter       string
	expectResult int
}

func TestMain(t *testing.T) {

	v := []values{
		{
			letter:       "I",
			expectResult: 1,
		},
		{
			letter:       "III",
			expectResult: 3,
		},
		{
			letter:       "IV",
			expectResult: 4,
		},
		{
			letter:       "V",
			expectResult: 5,
		},
		{
			letter:       "IX",
			expectResult: 9,
		},
		{
			letter:       "LVIII",
			expectResult: 58,
		},
		{
			letter:       "MCMXCIV",
			expectResult: 1994,
		},
	}

	for _, i := range v {
		val := romanToInt(i.letter)
		if val != i.expectResult {
			t.FailNow()
		}
	}
}
