package utils

import (
	"strconv"
	"time"
)

// GenerateKey is a function that generates a unique key for the user based on the current timestamp.
func GenerateKey() string {
	// unix nano is the current timestamp in nanoseconds

	// 	The strconv.FormatInt() function in Go converts an integer to its string representation. The function signature is:

	// func FormatInt(i int64, base int) string

	// i: The integer value you want to convert to a string.
	// base: The base for conversion, which can range from 2 to 36. For example, passing 10 would convert the integer to its decimal string representation, while passing 16 would convert it to hexadecimal.

	key := strconv.FormatInt(time.Now().UnixNano(),10)
	return key
}