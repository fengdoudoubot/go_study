package main

import(
	"errors"
	"fmt"
)

func f1(arg int) (int , error)  {
	if arg == 42 {
		return -1, errors.New("err 42!")
	}

	return arg + 3, nil
}

type argError struct{
	arg int 
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s",e.arg , e.prob)
}

