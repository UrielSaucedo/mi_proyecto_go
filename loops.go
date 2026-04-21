package main

import (
	"errors"
	"fmt"
)

func For() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func While() {
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func Range() {

	fmt.Println()
	names := []string{"Uriel", "Lucia", "Robert"}

	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println()
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	for key := range m {
		fmt.Println("key:", key)
	}

	fmt.Println()
	array := []int{10, 20, 30}

	sum := 0
	for _, value := range array {
		sum += value
	}
	fmt.Println("sum:", sum)

	fmt.Println("multiples variables")
	a := []int{1, 2, 3, 4, 5}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println("reversed:", a)
}

// ejemplo de función que puede fallar, para mostrar el manejo de errores u utilizar el If(me ayudo chatgpt a hacerlo y lo entendi)
func mightFail(value int) (string, error) {
	if value < 0 {
		return "", errors.New("negative value: input must be >= 0")
	}

	result := fmt.Sprintf("success: value %d is valid", value)
	return result, nil
}
