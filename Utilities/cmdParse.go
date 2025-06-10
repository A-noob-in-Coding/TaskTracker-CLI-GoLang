package Utilities

import (
	"strconv"
)

func checkNumber(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return true
	}
	return false

}

func Sanitize(args []string, decision *int) bool {
	lengthArgs := len(args)
	if lengthArgs < 2 {
		return false // not enough args to even check commands
	}

	if args[1] == "list" && lengthArgs == 2 {
		*decision = 1
		return true
	}
	if args[1] == "list-in-progress" && lengthArgs == 2 {
		*decision = 7
		return true
	}
	if args[1] == "list-done" && lengthArgs == 2 {
		*decision = 8
		return true
	}
	if args[1] == "mark-done" && lengthArgs == 3 && checkNumber(args[2]) {
		*decision = 6
		return true
	}
	if args[1] == "list-todo" && lengthArgs == 2 {
		*decision = 9
		return true
	}

	if args[1] == "mark-inprogress" && lengthArgs == 3 && checkNumber(args[2]) {
		*decision = 5
		return true
	}
	if args[1] == "update" && lengthArgs == 4 && checkNumber(args[2]) {
		*decision = 3
		return true
	}
	if args[1] == "add" && lengthArgs == 3 {
		*decision = 4
		return true
	}
	if args[1] == "delete" && lengthArgs == 3 && checkNumber(args[2]) {
		*decision = 2
		return true
	}

	return false
}
