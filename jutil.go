package main

import "fmt"

// printType prints the type of the variable passed in
func printType(variable interface{}) {
	switch variable.(type) {
	default:
		fmt.Println("Type unknown!")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case nil:
		fmt.Println("nil")
	}
}
