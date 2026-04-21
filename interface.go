package main

// Interface tiene que terminar en Er
type Stringer interface {
	String() string
}

// clase
type nombre struct {
	name string
}

// Metodo -(Esta funcion es un metodo de la clase user) - user se converitra en texto
func (n nombre) String() string {
	return n.name
}
