package tips

import "fmt"

func testLexicalScope() int {
	num1 := 0

	fmt.Println(num1)

	return num1
}
