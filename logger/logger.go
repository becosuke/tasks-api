package logger

import (
	"fmt"
	"log"
)

const (
	InvalidRequest = "invalid request"
)

func Print(args ...interface{}) {
	log.Print(args...)
}

func Error(err error) {
	log.Printf("%q", fmt.Sprintf("%+v", err))
}
