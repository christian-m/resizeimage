package main

import "fmt"

type errorList struct {
	errors []error
}

func (e *errorList) add(err error) {
	if err != nil {
		e.errors = append(e.errors, err)
	}
}

func (e *errorList) hasError() bool {
	return len(e.errors) > 0
}

func (e *errorList) Error() string {
	if !e.hasError() {
		return ""
	}
	out := fmt.Sprintf("Number of errors: %d", len(e.errors))
	for i, err := range e.errors {
		out = fmt.Sprintf("%s\n%d: %s", out, i+1, err.Error())
	}
	return out
}
