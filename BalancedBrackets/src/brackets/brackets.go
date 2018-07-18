package brackets

const openingBrace = '('
const closingBrace = ')'

// IsBalanced checks if given string is balanced in open and closed brackets
func IsBalanced(str string) bool{
	stack := []byte{}
	for _, s:= range str {
		if s == openingBrace {
			stack = append(stack, byte(s) )
		} else if s == closingBrace {
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if p != openingBrace {
					return false
				}
			} else {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}