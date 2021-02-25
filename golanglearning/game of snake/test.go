package main

import "fmt"

func main() {
	type ErrorCode int

	const (
		ERROR_SUCCESS ErrorCode = iota
		ERROR_FIRST
		ERROR_SECOND
		ERROR_THIRD
	)

	error_code := ERROR_SUCCESS
	fmt.Println("default: ", error_code) // default:  0

	error_code = ERROR_THIRD
	fmt.Println("Second: ", error_code) // Second:  2
}
