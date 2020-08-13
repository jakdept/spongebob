package main

import "os"

// lol tests?

func ExampleLowerCase() {
	os.Args = []string{"spongebob", "lower", "case", "test"}
	main()
	// Output: lOwEr cAsE TeSt
}

func ExampleUpperCase() {
	os.Args = []string{"spongebob", "UPPER", "CASE", "TEST"}
	main()
	// Output: lOwEr cAsE TeSt
}
