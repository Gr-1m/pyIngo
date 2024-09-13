package errnil

import (
	// "errors"
	"log"
)

func ErrNotnilLog(err error) {
	if err != nil {
		log.Panic(err)
	}
}
