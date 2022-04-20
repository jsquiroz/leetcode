package main

type Stack []rune

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}

func (s *Stack) Push(d rune) {
	*s = append(*s, d)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

var characteres = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func validParentheses(s string) bool {
	st := Stack{}
	for i := 0; i < len(s); i++ {
		if find(rune(s[i])) {
			st.Push(characteres[rune(s[i])])
		} else {
			val, _ := st.Pop()
			if val != rune(s[i]) {
				return false
			}
		}
	}

	return st.IsEmpty()
}

func find(r rune) bool {
	_, ok := characteres[r]

	return ok
}
