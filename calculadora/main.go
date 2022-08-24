/// ESTE ES UN EJERCICIO SIMPLE DE UNA CALCULADORA.
/// Curso Práctico de Go: Creación de un Servidor Web
/// video [6-9]
//  ? https://platzi.com/clases/1846-programacion-golang-2020/26752-descripcion-del-proyecto-1-calculadora/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	operacion := leerEntrada()
	operador := leerEntrada()
	fmt.Println(operacion)

	valores := strings.Split(operacion, operador) // Dividor de string
	fmt.Println(valores)

	switch operador {
	case "+":
		fmt.Println("Addition")
		operador1 := textToInt(valores[0])
		operador2 := textToInt(valores[1])
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println("Subtraction")
		operador1 := textToInt(valores[0])
		operador2 := textToInt(valores[1])
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println("Multiplication")
		operador1 := textToInt(valores[0])
		operador2 := textToInt(valores[1])
		fmt.Println(operador1 * operador2)
	case "/":
		fmt.Println("Division")
		operador1 := textToInt(valores[0])
		operador2 := textToInt(valores[1])
		fmt.Println(operador1 / operador2)
	default:
		fmt.Println(operador, "Operador no soportado")
	}

}

func textToInt(valor string) int {
	operador, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println("No es posible operar:", valor)
		return 0
	} else {
		return operador
	}
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
