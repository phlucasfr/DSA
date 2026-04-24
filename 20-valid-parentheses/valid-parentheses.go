func isValid(s string) bool {
    if len(s)%1 != 0 {
        return false
    }

    var stack = []byte{}
    closePatterns := map[byte]byte{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    for i := range s {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
        	stack = append(stack, s[i])
        } else {
            if len(stack) == 0 {
                return false
            } 

            if closePatterns[stack[len(stack)-1]] != s[i] {
                return false
            }

            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0    
}