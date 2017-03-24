package errors

import (
	"fmt"
	"log"
	"net/http"
)

func CheckAndLogError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}
