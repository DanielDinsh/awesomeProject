package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("my error to test package installation")
	fmt.Println(err.Error())
}
