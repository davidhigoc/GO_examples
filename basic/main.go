// ! 1. constantes, variables e impresion en consola
// ! 2. Operadores aritmeticos
// ! 3. Variables primitiva
// ! 4. Bucles o ciclos
// ! 5. Condicionales
// ! 6. KeyWords - [ defer - break - continue ]
// ! 7. Arrays - Slices
// ! 8. Mapas -> { "llave": "valor"}
// ! 9. Clases - Structs
// ! 10. Punteros, Punteros en Structs y Outputs personalizables de Struct
// ! 11. Interfaces y lista de Interfases
// ! 12. GoRoutines

package main

import (
	"custontext"
	"fmt"
	hl "go_pru1/basic/hola"
	"sync"
)

func main() {
	fmt.Printf("\n")
	fmt.Println("Hola mundo, llegue yo... ")
	fmt.Printf("\n")
	// ! 1. CONSTANTES Y VARIABLES
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
	// ! 2. OPERADORES ARITMETICOS
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
	// ! 3. VARIABLES PRIMITIVAS
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

	// ! 4. BUCLES Y CICLOS
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

	// ! 5. CONDICIONALES
	// Condicional IF
	var myBool bool = true
	fmt.Println("Variable bool =>", myBool)
	fmt.Println("Variable bool negada =>", !myBool)
	if !myBool {
		fmt.Printf("La variable es verdadera")
	} else {
		fmt.Println("La variable sigue siendo verdadera.")
		myBool = false
	}

	// Condicional SWITCH

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	switch valorInt := 200; valorInt {
	case 100:
		fmt.Println("Valor es 100")
	case 200:
		fmt.Println("Valor es 200")
	default:
		fmt.Println("No cumple condiciones")
	}

	fmt.Print("\n\n")

	// ! 6. Keywords
	defer fmt.Println("un DEFER cualquiera... Fin")
	fmt.Println("KeyWord DEFER - Activada")
	// ? break - continue -> Control de bucles

	fmt.Print("\n\n")

	// ! 7. Arrays - Slices
	// * Declaración
	// Array
	var array [4]int // Declaración por defecto de enteros - 0
	fmt.Println("Array por defecto: ", array)
	array[1] = 1
	array[2] = 2
	fmt.Println("Longitud:", len(array), "Capacidad max:", cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("Nuevo Slice -", slice, len(slice), cap(slice))
	// Metodos para el Slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4]) // ? Inclusivo - Exclusivo
	fmt.Println(slice[4:])
	// Append
	slice = append(slice, 20)
	fmt.Println("Se añade 20:", slice)
	// Append con nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println("Agragando un slice a otro slice:", slice)

	// Recorrido de Array y Slices
	slice1 := []string{"Hola", "Que", "Hace"}

	for i, valor := range slice1 { // _ para escapar - o solo se obvia el valor si no se necesita
		fmt.Println("Indice:", i, "Valor:", valor)
	}

	isPalidromo("ana") // Función para reconocer palabras palindromas [roma]

	fmt.Print("\n\n")

	// ! 8. Mapas
	m := make(map[string]int) // Definición
	m["Josecito"] = 14
	m["Pepo"] = 15
	fmt.Println(m)
	fmt.Println(m["Pepo"])

	// Recorrer map
	for i, v := range m {
		fmt.Println("Dato ", i, v)
	}

	// Encontrar valor y reconocer su existencia
	valor, ok := m["Josse"]
	fmt.Println("Existe llave?", ok, valor)

	fmt.Print("\n\n")

	// ! 9. Structs - class
	// Struct local
	var myCar car = car{marca: "Toyota", modelo: 2022} // STRUCT LOCAL
	fmt.Println(myCar)

	// Struct en otro archivo del proyecto
	var holis hl.Hola = hl.Hola{Saludo: "Hola", Despedida: "Adios"} // STRUCT PROJECT
	fmt.Println(holis)
	hl.Impresion("Uso de struct en diferente archivo del proyecto - Función")

	// Modulo en src - GOROOT - Package [Fabricación propia]
	fmt.Println("custontext: ", custontext.WhoSaidTea)
	var salu custontext.MensajePublicoModulo = custontext.MensajePublicoModulo{Mensaje: "Buen día", Dia: 23, Mes: 8, Año: 2022}
	fmt.Println(salu)

	fmt.Print("\n\n")

	// ! 10. Punteros [ & accede a direccion --- * accede a valor]
	// ! 10. Punteros, Punteros en Structs y Outputs personalizables de Struct

	aa := 50  // Valor
	bb := &aa // Variable que guarda la dirección en memoria de aa
	fmt.Println("Valor:", aa, "Dirección", bb, "ValorDir", *bb)
	// Cambiando el valor de aa
	*bb = 100
	fmt.Println("aa Cambiado:", aa)
	// Se crea un Struct - PC
	myPc := pc{ram: 16, disk: 200, marca: "msi"}
	myPc2 := pc{ram: 8, disk: 100, marca: "King"}
	fmt.Println("Normal: ", myPc)
	fmt.Print("\n")
	// ? AGREGAMOS UNA FUNCION A UN STRUCT
	myPc.ping() // Usamos una función de un struct
	myPc.duplicateRam()
	myPc2.duplicateRam()
	fmt.Println("Dupli x2", myPc)
	myPc.duplicateRam()
	fmt.Println("Dupli x4", myPc)
	fmt.Println("Dupli x2", myPc2)
	// ? OUTPUTS DE STRUCT
	// Se crea una función que devuelva un struct [customizado]
	fmt.Println(myPc.String())  // Salida 1
	fmt.Println(myPc2.String()) // Salida 2

	fmt.Print("\n\n")

	// ! 11. Interfaces y lista de Interfases
	myCuadrado := cuadrado{base: 2}
	fmt.Println("Cuadrado:", myCuadrado)
	myRectangulo := rectangulo{base: 2, altura: 4}
	fmt.Println("Rectangulo:", myRectangulo)
	calculate(myCuadrado)
	calculate(myRectangulo)
	// lista de interfaces
	myInterfaces := []interface{}{1, "hola", 3.1416}
	fmt.Println(myInterfaces...)

	fmt.Print("\n\n")

	// ! 12. GoRoutines
	var wg sync.WaitGroup

	// say("Hola mundo")   // consumo funcion normalmente
	wg.Add(1)
	go say("Ya llegue", &wg) // consumo de función con GoRoutine
	// time.Sleep(time.Millisecond * 500)
	wg.Wait()
}

///
///
///
///
///
///
///

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Salida say:", text)
}

///
///

type figuras2D interface { // INTERFACE
	area() float64
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calculate(f figuras2D) { //
	fmt.Println("Area:", f.area())
}

///
///

type pc struct {
	ram   int
	disk  int
	marca string
}

// FUNCIÓN PARA PC
func (elPc pc) ping() {
	fmt.Println(elPc.marca, "Pong - [METODO - FUNC]")
}
func (elPc *pc) duplicateRam() { // Accedemos a los valores del struct que me entregan en la func
	elPc.ram = elPc.ram * 2
}
func (elPc pc) String() string { // SALIDA CUSTOMIZABLE DE UN STRUCT
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de MEMORIA y es un %s", elPc.ram, elPc.disk, elPc.marca)
}

///
///

type car struct {
	marca  string
	modelo int
}

///
///

func isPalidromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println(text, "es palindromo")
	} else {
		fmt.Println(text, "no es palindromo")
	}
}

///
///

func bucles() {
	fmt.Printf(".")
}
