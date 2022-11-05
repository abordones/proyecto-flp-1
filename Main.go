package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	connectionEstablished := conectionBD()

	//Son inserciones previsionales -> Funcionan pero faltan modificaciones.
	//Al estar pendiente la confeccion y concretacion de la BD -> todas las inserciones carecen de FK.

	insertQuestion(connectionEstablished)
	insertPatient(connectionEstablished)
	insertUser(connectionEstablished)
	insertAnswers(connectionEstablished)

	//7.-Cerrar y finalizar.
	fmt.Println("FIN.")
}

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

func insertQuestion(connectionEstablished *sql.DB) {

	fmt.Print("Formula tu pregunta: ")
	q := bufio.NewReader(os.Stdin)
	question, _ := q.ReadString('\n')

	fmt.Print("Escribe su descripcion: ")
	d := bufio.NewReader(os.Stdin)
	description, _ := d.ReadString('\n')

	InsertQuestion, err := connectionEstablished.Prepare("INSERT INTO questions (ID, QUESTION, DESCRIPTION_Q, DELETE_AT, TEST_ID) VALUES (NULL, ?, ?, 1, NULL); ")
	if err != nil {
		panic(err.Error())
	}
	InsertQuestion.Exec(question, description)

	fmt.Println("Pregunta ingresada con exito.")
}

func insertPatient(connectionEstablished *sql.DB) {

	var RUN int
	fmt.Print("Ingresa RUN del paciente: ")
	fmt.Scanln(&RUN)

	var DV string
	fmt.Print("Ingresa digito verificador del paciente: ")
	fmt.Scanln(&DV)

	fmt.Print("Ingresa nombre(s) del paciente: ")
	n := bufio.NewReader(os.Stdin)
	names, _ := n.ReadString('\n')

	var fatherName string
	fmt.Print("Ingresa apellido paterno del paciente: ")
	fmt.Scanln(&fatherName)

	var motherName string
	fmt.Print("Ingresa apellido materno del paciente: ")
	fmt.Scanln(&motherName)

	var phone int
	fmt.Print("Ingresa el numero de contacto del paciente: ")
	fmt.Scanln(&phone)

	var email string
	fmt.Print("Ingresa el correo electronico del paciente: ")
	fmt.Scanln(&email)

	fmt.Print("Ingresa una observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	insertPatient, err := connectionEstablished.Prepare("INSERT INTO patients (ID, RUN, DV, NAMES, FATHERNAME, MOTHERNAME, PHONE, MAIL, OBSERVATION, ACTIVE) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, '1')")
	if err != nil {
		panic(err.Error())
	}
	insertPatient.Exec(RUN, DV, names, fatherName, motherName, phone, email, observation)

	fmt.Println("Paciente ingresado con exito.")
}

func insertUser(connectionEstablished *sql.DB) {

	var RUN int
	fmt.Print("Ingresa tu RUN: ")
	fmt.Scanln(&RUN)

	var DV string
	fmt.Print("Ingresa tu digito verificador: ")
	fmt.Scanln(&DV)

	fmt.Print("Ingresa tu(s) nombre(s): ")
	n := bufio.NewReader(os.Stdin)
	names, _ := n.ReadString('\n')

	var fatherName string
	fmt.Print("Ingresa tu apellido paterno: ")
	fmt.Scanln(&fatherName)

	var motherName string
	fmt.Print("Ingresa tu apellido materno: ")
	fmt.Scanln(&motherName)

	var email string
	fmt.Print("Ingresa el correo electronico del paciente: ")
	fmt.Scanln(&email)

	var password string
	fmt.Print("Ingresa tu contrasena: ")
	fmt.Scanln(&password)

	insertUser, err := connectionEstablished.Prepare("INSERT INTO users (ID, RUN, DV, NAMES, FATHERNAME, MOTHERNAME, EMAIL, PASSWORD, ACTIVE) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, '1');")
	if err != nil {
		panic(err.Error())
	}
	insertUser.Exec(RUN, DV, names, fatherName, motherName, email, password)

	fmt.Println("Usuario ingresado con exito.")

}

func insertAnswers(connectionEstablished *sql.DB) {

	var point int
	fmt.Print("Ingresa el puntaje (entre 0 y 3): ")
	fmt.Scanln(&point)

	fmt.Print("Ingresa una observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	insertAnswers, err := connectionEstablished.Prepare("INSERT INTO answers (ID, POINT, OBSERVATION) VALUES (NULL, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertAnswers.Exec(point, observation)

	fmt.Println("Respuesta ingresada.")

}
