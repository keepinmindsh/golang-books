package basic

import (
	"fmt"
	"strings"
)

/**

function이 return된 뒤에 실행될 것을 정의하기 위해 사용된다.
예를 들어 func가 return된 이후에 파일을 삭제한다던가 api로 요청을 보낸다던가 하는 행위들을 defer 를 통해 정의할 수 있다.

*/
func DeferSample(name string) (length int, uppercase string) {
	defer fmt.Println("im Done!!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
