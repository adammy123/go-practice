package main

import (
	"fmt"

	"adam.com/toto/generateparentheses/gp"
)

const N = 3

func main() {
	for i := 1; i <= N; i++ {
		fmt.Println(i, gp.GenerateParentheses(i))
	}
}
