package tips

import "time"

/*

구조체도 함수와 마찬가지로 필드명을 사용하여 외부에
노출시킬 필드(Public)와 노출 시키지 않을 필드(Private)를
설정할 수 있습니다. 필드명이 대문자로 시작하면 외부에서 사용이 가능하지만,
소문자로 시작하는 경우에는 외부에서 사용이 불가능하다.

*/

type StuructureSample struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type ClassInfo struct {
	Class int
	No    int
}

type Student struct {
	ClassInfo
	Name  string
	Class int
	No    int
}
