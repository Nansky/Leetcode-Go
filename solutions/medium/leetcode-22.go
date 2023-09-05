// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// Example 1:
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

// Example 2:
// Input: n = 1
// Output: ["()"]

// Constraints:
// 1 <= n <= 8

func generateParenthesis(n int) []string {

	vp := make([]string, 0)
	stack := make([]string, 0)

	var dfs func(leftBracket, rightBracket int)

	dfs = func(leftBracket, rightBracket int) {
		if leftBracket == n && rightBracket == n && leftBracket == rightBracket {
			vp = append(vp, strings.Join(stack, ""))
			return
		}

		if leftBracket < n {
			stack = append(stack, "(")
			// fmt.Println(leftBracket)
			dfs(leftBracket+1, rightBracket)
			pop(&stack)
		}

		if rightBracket < leftBracket {
			stack = append(stack, ")")
			// rightBracket += 1
			dfs(leftBracket, rightBracket+1)
			pop(&stack)
		}
	}

	dfs(0, 0)

	return vp
}

func pop(stack *[]string) {
	n := len(*stack)

	*stack = (*stack)[:n-1]
}
