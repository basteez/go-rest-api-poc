package utils

import "log"

func LogOnError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func PanicOnError(err error, message string) {
	if err != nil {
		panic(message)
	}
}
