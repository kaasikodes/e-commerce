package utils

import "fmt"

func Recover() {
	if err := recover(); err != nil {
		fmt.Println("From recover utils")
		fmt.Println(err)
	}
}