package brackets

func isOpeningBrace(ch byte) bool {
    return ch == '{' || ch == '[' || ch == '('
}

func isClosingBrace(ch byte) bool {
    return ch == '}' || ch == ']' || ch == ')'
}

func getClosingPair(ch byte) byte {
    switch ch {
        case '{': return '}'
        case '[': return ']'
        case '(': return ')'
    }    
    return ' '
}

// Complete the isBalanced function below.
func IsBalanced(str string) bool {
    stack := []byte{}
	for _, s:= range str {
        if isOpeningBrace(byte(s)) {
			stack = append(stack, byte(s) )
        } else if isClosingBrace(byte(s)) {
			if len(stack) > 0 {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
                if !isOpeningBrace(byte(p)) || byte(s) != getClosingPair(p) {
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
