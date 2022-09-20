package tips

import "fmt"

func TypeSample() {
	x := 1

	//x = "String"

	fmt.Println(x)

	q, w, r := returnArgs()

	fmt.Println(q, w, r)
}

func returnArgs() (int, int, int) {
	return 1, 3, 4
}

func defaultArg() {

}