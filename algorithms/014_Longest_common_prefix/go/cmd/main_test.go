package main

import (
	"log"
	"testing"
)

type values struct {
	arr          []string
	expectResult string
}

func TestLongestCommonPrefix(t *testing.T) {

	v := []values{
		{
			arr:          []string{"flower", "flow", "fligth"},
			expectResult: "fl",
		},
		{
			arr:          []string{"dog", "racecar", "car"},
			expectResult: "",
		},
		{
			arr:          []string{"a"},
			expectResult: "a",
		},
		{
			arr:          []string{"ab", "a"},
			expectResult: "a",
		},
		{
			arr:          []string{"cir", "car"},
			expectResult: "c",
		},
		{
			arr:          []string{"reflower", "flow", "flight"},
			expectResult: "",
		},
	}

	for _, i := range v {
		val := longestCommonPrefix(i.arr)
		if val != i.expectResult {
			log.Println(i.arr, val, i.expectResult)
			t.FailNow()
		}
	}

}
