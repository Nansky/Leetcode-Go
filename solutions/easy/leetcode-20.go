// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Constraints:
// s consists of parentheses only '()[]{}'.

func isValid(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}

	st := []rune{}
	m := make(map[rune]rune)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	for _, sr := range s {
		if _, exists := m[sr]; exists {
			st = append(st, sr)
		} else if len(st) == 0 || m[st[len(st)-1]] != sr {
			return false
		} else {
			st = st[:len(st)-1]
		}
	}

	return len(st) == 0
}