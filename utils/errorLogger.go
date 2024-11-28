package utils

import "log"

func LogOnError(err error, message string) {
	if err != nil {
		log.Println("[ERROR] ", message, ": ", err)
	}
}

func PanicOnError(err error, message string) {
	if err != nil {
		panic(message)
	}
}
