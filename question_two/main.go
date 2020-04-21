package main

import (
	"github.com/bmsandoval/question_two/entry"

	// this is required for self-registration to work properly
	// complicated import process allows this
	_ "github.com/bmsandoval/question_two/formats"
)

// always nice to have a clean main :D
func main() {
	entry.Entry()
}