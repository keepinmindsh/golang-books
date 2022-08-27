package problems

import (
	"fmt"
	"strconv"
)

func PrintSong(printCount int) {

	for i := 0; i < printCount; i++ {
		multiply := 10 * i

		fmt.Printf("타잔이  %s 원 짜리 팬티를 입고, %s 원 짜리 칼을 차고 노래를 한다. \r\n", strconv.Itoa(multiply), strconv.Itoa(multiply+10))
	}
}
