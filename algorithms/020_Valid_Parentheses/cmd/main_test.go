package main

import (
	"log"
	"testing"
)

type values struct {
	arr          string
	expectResult bool
}

func TestValidParenthesis(t *testing.T) {

	v := []values{
		{
			arr:          "[]{}()",
			expectResult: true,
		},
		{
			arr:          "[)",
			expectResult: false,
		},
		{
			arr:          "(){}}{",
			expectResult: false,
		},
	}

	for _, i := range v {
		val := validParentheses(i.arr)
		if val != i.expectResult {
			log.Println(i.arr, val, i.expectResult)
			t.FailNow()
		}
	}

}
