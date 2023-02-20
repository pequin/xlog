package xlog

import (
	"log"
)

func isNotNil(v ...any) bool {

	for i := 0; i < len(v); i++ {
		if v[i] == nil {
			return false
		}
	}
	return true
}

// If values is not nil followed call to log.Fatalf().
func Fatallf(format string, v ...any) {

	if isNotNil(v...) {
		log.Fatalf(format, v...)
	}
}

// If values is not nil followed call to log.Fatalln().
func Fatalln(v ...any) {

	if isNotNil(v...) {
		log.Fatalln(v...)
	}
}
