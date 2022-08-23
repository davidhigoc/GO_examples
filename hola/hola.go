package hola

import "fmt"

type Hola struct {
	Saludo    string
	Despedida string
}

func Impresion(mess string) {
	fmt.Println("Mensaje:", mess)
}
