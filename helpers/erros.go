package helpers

import "log"

func IsPanicErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}
