package utils

import "log"

type MyError struct {
	function string
	err      error
}

func CheckError(function string, err error) {
	if err != nil {
		myerror := &MyError{function, err}
		log.Fatal(myerror)
		panic(myerror)
	}
}
