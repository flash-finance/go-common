package util

import "fmt"

func CatchError() {
	if err := recover(); err != nil {
		fmt.Errorf("recover error:%v\n", err)
	}
}
