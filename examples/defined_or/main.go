package main

import "github.com/rokuosan/vars"

func main() {
	// Example usage of DefinedOr function
	type testStruct struct {
		Field string
	}

	var nilPtr *testStruct
	defaultVal := testStruct{"default"}

	result := vars.DefinedOr(nilPtr, defaultVal)
	println(result.Field) // Output: default

	value := testStruct{"actual"}
	ptr := &value

	result = vars.DefinedOr(ptr, defaultVal)
	println(result.Field) // Output: actual
}
