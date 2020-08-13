package main

import "os"

func ExampleMain_lowercase() {
	os.Args = []string{"spongebob", "lower", "case", "test"}
	main()
	// Output: lOwEr cAsE TeSt
}

func ExampleMain_uppercase() {
	os.Args = []string{"spongebob", "UPPER", "CASE", "TEST"}
	main()
	// Output: lOwEr cAsE TeSt
}
