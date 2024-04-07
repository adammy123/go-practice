package gp

//Generate all combinations of well-formed parentheses.
//A parenthesis combination is well-formed if each opening parenthesis "("
// is closed by a matching closing parenthesis ")" in the correct order.

// For example, "()", "()()", and "((()))" are all combinations of well-formed parentheses,
// while ")(", "((" and "(()" are not.
func GenerateParentheses(n int) []string {
	res := []string{}
	addToRes(n, n, res)
	return res
}

func addToRes(numOpen, numClose int, res []string)

// add '(' if still avail
// add ')' if sttill avail && added '('s > ')'
