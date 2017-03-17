package errors

import "log"

func CheckAndLogError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
