package main

import (
	"fmt"
)

func main() {
	fmt.Println("TITULO")
	fmt.Println("Inicio del cuestionario")
	//1.-Imprimir Instrucicciones.
	/*--- Extraer informacion en BD.--- */

	//2.-Recibir al participante.
	var day int
	fmt.Print("Ingresar dia: ")
	fmt.Scanln(&day)

	var month int
	fmt.Print("Ingresar mes: ")
	fmt.Scanln(&month)

	var year int
	fmt.Print("Ingresar anio: ")
	fmt.Scanln(&year)

	var rut string
	fmt.Print("Ingresar rut: ")
	fmt.Scanln(&rut)

	var name string
	fmt.Print("Ingresar nombre: ")
	fmt.Scanln(&name)

	var lastName string
	fmt.Print("Ingresar apellido: ")
	fmt.Scanln(&lastName)

	var age int
	fmt.Print("Ingresar edad: ")
	fmt.Scanln(&age)
	/*--- Guardar informacion en BD.--- */

	//3.-Realizar preguntas.
	/*--- Extraer informacion de BD.--- */

	//4.-Obtener respuestas.
	var answers [10]int
	var num int
	var i int

	for i = 0; i < 10; i++ { //relleno
		fmt.Printf("Respuesta de pregunta (%v): ", i+1)
		fmt.Scanln(&num)
		if num >= 1 && num <= 3 {
			answers[i] = num
		} else {
			i--
		}
	}
	/*--- Guardar informacion en BD.--- */

	//5.-Calcular resultados.
	var result int
	for i = 0; i < 10; i++ {
		result = result + answers[i]
	}
	/*--- Guardar informacion en BD.--- */

	//6.-Mostrar resultados.
	fmt.Printf("El resultado es: %v\n", result) //Respuesta arbitraria.
	if result < 10 {
		fmt.Printf("Participante %v %v se encuentra en buen estado.\n", name, lastName)
	}
	if result >= 10 && result <= 20 {
		fmt.Printf("Participante %v %v se encuentra en condicion estable.\n", name, lastName)
	}
	if result > 20 && result <= 25 {
		fmt.Printf("Participante %v %v requiere atencion.\n", name, lastName)
	}
	if result > 25 {
		fmt.Printf("Participante %v %v se le diagnostica depresion.\n", name, lastName)
	}

	//7.-Cerrar y finalizar.
	fmt.Println("FIN.")
}
