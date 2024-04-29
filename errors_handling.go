package main

import (
	"errors"
	"fmt"
)

// Errors
func main_errors() {

	q, err := divide(20, 0)
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println(q)

	err1 := errors.New("original error")

	wrappedErr := fmt.Errorf("Wrapped Error : %s", err1)

	fmt.Println(wrappedErr)

}

func divide(x, y int) (int, error) {

	if y == 0 {
		return 0, &MyError{"Something went Wrong", 504}
	}
	return x / y, nil
}

type MyError struct {
	Msg     string
	ErrCode int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error :%d : %s", e.ErrCode, e.Msg)
}
