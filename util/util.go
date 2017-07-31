package util

import (
	"log"
)

// Check is a function for error checking
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
