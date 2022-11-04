package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func conectionBD() (conection *sql.DB) {

	driver := "mysql"
	user := "root"
	password := ""
	nameBD := "bd-flp-1"

	conection, err := sql.Open(driver, user+":"+password+"@tcp(127.0.0.1)/"+nameBD)
	if err != nil {
		panic(err.Error())
	}
	return conection
}

func main() {

	connectionEstablished := conectionBD()

	insertQuestion(connectionEstablished)

	/*var question string
	//No tolera espacios, acepta_
	fmt.Print("Formula tu pregunta: ")
	fmt.Scanln(&question)

	//InsertUser, err := connectionEstablished.Prepare("INSERT INTO usuarios (nombre, edad) VALUES ('Charles', '33') ")
	InsertUser, err := connectionEstablished.Prepare("INSERT INTO questions (ID, QUESTION, DESCRIPTION_Q, DELETE_AT, TEST_ID) VALUES (NULL, ?, 'Conocer patron de sueño', NULL, NULL); ")

	if err != nil {
		panic(err.Error())
	}
	InsertUser.Exec(question)*/

	/*
		fmt.Println("TITULO")
		fmt.Println("Inicio del cuestionario")
		//1.-Imprimir Instrucicciones.
		/*--- Extraer informacion en BD.---

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
	/*--- Extraer informacion de BD.---

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
	/*--- Guardar informacion en BD.---

	//5.-Calcular resultados.
	var result int
	for i = 0; i < 10; i++ {
		result = result + answers[i]
	}
	/*--- Guardar informacion en BD.---

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
	*/
}

func insertQuestion(connectionEstablished *sql.DB) {

	//No tolera espacios, acepta_
	var question string
	fmt.Print("Formula tu pregunta: ")
	fmt.Scanln(&question)

	var description string
	fmt.Print("Escribe su descripcion: ")
	fmt.Scanln(&description)

	InsertUser, err := connectionEstablished.Prepare("INSERT INTO questions (ID, QUESTION, DESCRIPTION_Q, DELETE_AT, TEST_ID) VALUES (NULL, ?, ?, 1, NULL); ")

	if err != nil {
		panic(err.Error())
	}
	InsertUser.Exec(question, description)

	fmt.Println("Pregunta Ingresada con exito.")
}
