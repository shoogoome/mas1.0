package exception

import "fmt"

// 系统异常输出
func OutputException(msg string, err error) {
	panic(fmt.Sprintf("==> %s: %s", msg, err.Error()))
}

