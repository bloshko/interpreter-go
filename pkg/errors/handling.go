package errors

import (
	"errors"
	"fmt"
)

func Error(line int, message string) {
	err := errors.New(message)

	report(line, "", err)
}

func report(line int, where string, err error) {
	fmt.Println(err)
}
