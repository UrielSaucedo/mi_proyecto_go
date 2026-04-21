package main

import "fmt"

// new te da un valor por defecto = 0
func NewEjemplo() {
	p := new(int)

	fmt.Println("new int value:", *p) // 0

	*p = 10
	fmt.Println("after set:", *p)
}

// make se usa para inicializar slices, maps y channels, --ejemplo de slice-- (recordatorio, no entendi que es channels)...slice = array, map =	diccionario, channel = canal de comunicacion entre goroutines?
func MakeSliceEjemplo() {
	s := make([]int, 3)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println("slice:", s)
}

// comparacion de new y make
func Comparison() {
	// new slice (incorrecto) - deberia ser make, asi que devolvera un nil slice
	p := new([]int)

	fmt.Println("new([]int) pointer:", p)
	fmt.Println("value inside:", *p) // nil slice

	// make slice (correcto) - solo usar make para slices, maps y channels
	v := make([]int, 3)
	fmt.Println("make([]int):", v)
}
