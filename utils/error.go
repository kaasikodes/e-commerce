package utils

import "fmt"

func ErrHandler(err error) error {
	if err != nil {

		fmt.Println("From utils")
		fmt.Println(err)
		panic(err)
		// return err
	}
	return nil
}