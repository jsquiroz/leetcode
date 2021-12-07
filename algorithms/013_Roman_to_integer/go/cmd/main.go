package main

var letters = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if len(s) == 1 {
		return letters[string(s[0])]
	}

	res := letters[string(s[len(s)-1])]

	for i := len(s) - 2; i >= 0; i-- {
		res += resolve(letters[string(s[i])], letters[string(s[i+1])])
	}

	return res
}

func resolve(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		return a
	}

	return -a
}
