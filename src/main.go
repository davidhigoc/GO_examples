// ! 1. constantes, variables e impresion en consola
// ! 2. Operadores aritmeticos
// ! 3. Variables primitiva
// ! 4. Bucles o ciclos
// ! 5. Condicionales

package main

import "fmt"

func main() {
	fmt.Printf("\n")
	fmt.Println("Hola mundo, llegue yo... ")
	fmt.Printf("\n")
	// ! 1.
	/// EJEMPLO BASICO
	// ? DECLARACIÓN DE CONSTANTES

	const pi float64 = 3.1416
	const pi2 = pi
	fmt.Println(pi, pi2)

	// ? DECLARACION DE VARIABLES
	base := 12
	var altura int = 170
	var area int // default INT -> 0

	fmt.Println(base, altura, area)

	// zero values
	var a int
	var b float64
	var c string // default STRING -> "" [vacio]
	var d bool
	fmt.Println(a, b, c, d)

	// Calculo de area de un cuadrado [ base x altura]
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del Cuadrado:", areaCuadrado)

	fmt.Printf("\n")
	// ! 2.
	x := 10
	y := 50

	// ? Suma
	result := x + y
	fmt.Println("Suma de X y Y:", result)

	// ? Resta
	result = x - y
	fmt.Println("Resta de X y Y:", result)

	// ? Multiplicación
	result = x * y
	fmt.Println("Multiplicación de X y Y:", result)

	// ? Divición
	result = x / y
	fmt.Println("Divición de X y Y:", result)

	// ? Modulo
	result = x % y
	fmt.Println("Modulo de X y Y:", result)

	// ? Incremental
	x++
	fmt.Println("Incremental de X y Y:", x)

	// ? Decremental
	x--
	fmt.Println("Decremental de X y Y:", x)

	fmt.Printf("\n")
	// ! 3.
	//  Numeros enteros
	//  int = Depende del OS (32 o 64 bits)
	//  int8 = 8bits = -128 a 127
	//  int16 = 16bits = -2^15 a 2^15-1
	//  int32 = 32bits = -2^31 a 2^31-1
	//  int64 = 64bits = -2^63 a 2^63-1

	//  Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//  uint = Depende del OS (32 o 64 bits)
	//  uint8 = 8bits = 0 a 127
	//  uint16 = 16bits = 0 a 2^15-1
	//  uint32 = 32bits = 0 a 2^31-1
	//  uint64 = 64bits = 0 a 2^63-1

	//  numeros decimales
	//   float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	//   float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//  textos y booleanos
	//  string = ""
	//  bool = true or false

	//  numeros complejos
	//  Complex64 = Real e Imaginario float32
	//  Complex128 = Real e Imaginario float64
	//  Ejemplo : c:=10 + 8i

	// ! 4.
	//	For condicional
	for i := 0; i <= 10; i++ {
		fmt.Print("   ", i)
	}
	fmt.Printf("\n")
	// For tipo While
	counter := 10
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}

	// For tipo loop()
	counter = 100
	for counter != 0 {
		bucles()
		counter--
	}

	fmt.Print("\n\n")

	// ! 5.
}

func bucles() {
	fmt.Printf(".")
}
