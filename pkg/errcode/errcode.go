package pkg

import "fmt"

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprint("错误码%d已经存在，请换一个", code))

	}
	codes[code] = msg

}
