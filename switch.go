package main

import "fmt"

// Directamente me ayudo chatgpt a hacer estos ejemplos de switch.
// 🔀 switch básico
func SwitchBasic(x int) {
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}
}

// 🔥 switch sin condición (muy Go)
func SwitchNoCondition(x int) {
	switch {
	case x > 10:
		fmt.Println("greater than 10")
	case x > 5:
		fmt.Println("greater than 5")
	default:
		fmt.Println("small number")
	}
}

// 🔥 múltiples valores en un case
func SwitchMultiple(c byte) {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("should escape")
	default:
		fmt.Println("normal character")
	}
}

// 🔥 ejemplo tipo del documento (unhex)
func Unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
