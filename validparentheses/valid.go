package validparentheses

func isValid(s string) bool {

	leftstack := make([]rune, 0, len(s))
	runes := []rune(s)
	var mymap = map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		switch c {
		//判断是不是左括号
		case '(', '[', '{':
			//push it to the top of the stack
			leftstack = append(leftstack, c)
		//判断右括号
		case ')', ']', '}':
			if len(leftstack) == 0 || leftstack[len(leftstack)-1] != mymap[c] {
				return false
			}
			leftstack = leftstack[:len(leftstack)-1]
		}
	}
	return len(leftstack) == 0
}
