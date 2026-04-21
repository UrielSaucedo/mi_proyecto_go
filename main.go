package main

import (
	"fmt"
	"mi_proyecto_go/my_package"
)

func main() {
	fmt.Println("test")

	//esta funcion es parte del paquete my_package, por lo que es necesario importarla
	user := my_package.New("Uriel")
	fmt.Println(user.Name)

	// son parte del mismo paquete asi que no es necesario importar
	fmt.Println("-- Loops --")
	For()
	While()
	Range()

	// añadi en loops un ejemplo de error con un if
	fmt.Println("-- mightFail --")

	// caso exitoso
	result, err := mightFail(5)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("-- allocation --")

	NewEjemplo()
	MakeSliceEjemplo()
	Comparison()

	// caso error
	result, err = mightFail(-1)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(result)
	}

	// Interface no necesita ser importada porrque es parte del mismo paquete
	fmt.Println("-- Interfaces --")
	n := nombre{name: "Robert"}

	var s Stringer = n
	fmt.Println(s.String())

	// Switch, esto me explico y ayudo chatgpt a hacerlo. (lo entendi)

	fmt.Println("-- Switch basico --")
	SwitchBasic(1)
	SwitchBasic(3)

	fmt.Println("--Switch sin condicion --")
	SwitchNoCondition(7)
	SwitchNoCondition(12)

	SwitchMultiple('?')

	fmt.Println("unhex:", Unhex('A'))
}
