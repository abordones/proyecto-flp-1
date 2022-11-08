package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID         int
	RUN        int
	DV         string
	names      string
	fatherName string
	motherName string
	email      string
	password   string
	active     bool
}

type Patient struct {
	ID          int
	RUN         int
	DV          string
	names       string
	fatherName  string
	motherName  string
	phone       int
	email       string
	observation string
	active      bool
}

func main() {

	connectionEstablished := conectionBD()

	//Son inserciones previsionales -> Funcionan pero faltan modificaciones.
	//Al estar pendiente la confeccion y concretacion de la BD -> todas las inserciones carecen de FK.
	//insertUser(connectionEstablished)
	//insertPatient(connectionEstablished)
	insertTest(connectionEstablished)

	//insertQuestion(connectionEstablished, 1)
	//insertAnswer(connectionEstablished)

	//Reciben conexion y una 'ID'
	//disableUser(connectionEstablished, 1)    //No confirma si el usuario existe.
	//disablePatient(connectionEstablished, 1) //No confirma si el usuario existe.

	//readAllUsers(connectionEstablished)
	//readAllPatients(connectionEstablished)

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

func insertQuestion(connectionEstablished *sql.DB, ID_test int) {

	fmt.Print("Formula tu pregunta: ")
	q := bufio.NewReader(os.Stdin)
	question, _ := q.ReadString('\n')

	fmt.Print("Escribe su descripcion: ")
	d := bufio.NewReader(os.Stdin)
	description, _ := d.ReadString('\n')

	InsertQuestion, err := connectionEstablished.Prepare("INSERT INTO questions (ID_Q, ID_T, QUESTION_Q, DESCRIPTION_Q, ACTIVE_Q) VALUES (NULL, ?, ?, ?, 1); ")
	if err != nil {
		panic(err.Error())
	}
	InsertQuestion.Exec(ID_test, question, description)

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

	var birthday string
	fmt.Print("Ingresa tu fecha de nacimineto (yyyy-mm-dd): ")
	fmt.Scanln(&birthday)

	fmt.Print("Ingresa una observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	insertPatient, err := connectionEstablished.Prepare("INSERT INTO patients (ID_P, RUN_P, DV_P, NAME_P, FATHERNAME_P, MOTHERNAME_P, PHONE_P, EMAIL_P, BIRTHDAY_P, OBSERVATION_P, ACTIVE_P) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, '1')")
	if err != nil {
		panic(err.Error())
	}
	insertPatient.Exec(RUN, DV, names, fatherName, motherName, phone, email, birthday, observation)

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

	var birthday string
	fmt.Print("Ingresa tu fecha de nacimineto (yyyy-mm-dd): ")
	fmt.Scanln(&birthday)

	var email string
	fmt.Print("Ingresa tu correo electronico: ")
	fmt.Scanln(&email)

	var password string
	fmt.Print("Ingresa tu contrasena: ")
	fmt.Scanln(&password)

	insertUser, err := connectionEstablished.Prepare("INSERT INTO users (ID_U, RUN_U, DV_U, NAME_U, FATHERNAME_U, MOTHERNAME_U, BIRTHDAY_U, EMAIL_U, PASSWORD_U, ACTIVE_U) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, '1');")
	if err != nil {
		panic(err.Error())
	}
	insertUser.Exec(RUN, DV, names, fatherName, motherName, birthday, email, password)

	fmt.Println("Usuario ingresado con exito.")

}

func insertTest(connectionEstablished *sql.DB) {

	fmt.Print("Nombra al test: ")
	n := bufio.NewReader(os.Stdin)
	name, _ := n.ReadString('\n')

	var cutPoint int
	fmt.Print("Ingresa el puntaje de corte: ")
	fmt.Scanln(&cutPoint)

	var matchPoint int
	fmt.Print("Ingresa puntaje match: ")
	fmt.Scanln(&matchPoint)

	fmt.Print("Ingresa una observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	insertTest, err := connectionEstablished.Prepare("INSERT INTO tests (ID_T, NAME_T, CUTPOINT_T, MATCHPOINT_T, OBSERVATION_T, ACTIVE_T) VALUES (NULL, ?, ?, ?, ?, '1')")
	if err != nil {
		panic(err.Error())
	}
	insertTest.Exec(name, cutPoint, matchPoint, observation)

	fmt.Println("Test ingresado.")
}

func insertAnswer(connectionEstablished *sql.DB) {

	var point int
	fmt.Print("Ingresa el puntaje (entre 0 y 3): ")
	fmt.Scanln(&point)

	fmt.Print("Ingresa una observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	insertAnswer, err := connectionEstablished.Prepare("INSERT INTO answers (ID, POINT, OBSERVATION) VALUES (NULL, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertAnswer.Exec(point, observation)

	fmt.Println("Respuesta ingresada.")

}

func readAllUsers(connectionEstablished *sql.DB) {

	readUser, err := connectionEstablished.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	user := User{} //Variable := Struct
	arrayUser := []User{}

	for readUser.Next() {
		var ID int
		var RUN int
		var DV string
		var names string
		var fatherName string
		var motherName string
		var email string
		var password string
		var active bool
		err = readUser.Scan(&ID, &RUN, &DV, &names, &fatherName, &motherName, &email, &password, &active)
		if err != nil {
			panic(err.Error())
		}

		user.ID = ID
		user.RUN = RUN
		user.DV = DV
		user.names = names
		user.fatherName = fatherName
		user.motherName = motherName
		user.email = email
		user.password = password
		user.active = active
		arrayUser = append(arrayUser, user)

	}
	fmt.Println(arrayUser)
}

func readAllPatients(connectionEstablished *sql.DB) {

	readPatient, err := connectionEstablished.Query("SELECT * FROM patients")
	if err != nil {
		panic(err.Error())
	}

	patient := Patient{} //Variable := Struct
	arrayPatient := []Patient{}
	for readPatient.Next() {
		var ID int
		var RUN int
		var DV string
		var names string
		var fatherName string
		var motherName string
		var phone int
		var email string
		var observation string
		var active bool
		err = readPatient.Scan(&ID, &RUN, &DV, &names, &fatherName, &motherName, &phone, &email, &observation, &active)
		if err != nil {
			panic(err.Error())
		}

		patient.ID = ID
		patient.RUN = RUN
		patient.DV = DV
		patient.names = names
		patient.fatherName = fatherName
		patient.motherName = motherName
		patient.phone = phone
		patient.email = email
		patient.observation = observation
		patient.active = active
		arrayPatient = append(arrayPatient, patient)
	}
	fmt.Println(arrayPatient)

}

func disableUser(connectionEstablished *sql.DB, ID int) {

	idUser := ID
	disableUser, err := connectionEstablished.Prepare("UPDATE users SET ACTIVE=false WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	disableUser.Exec(idUser)

	fmt.Printf("El usuario ID (%v) ha sido deshabilitado.\n", idUser)
}

func disablePatient(connectionEstablished *sql.DB, ID int) {

	idPatient := ID
	disablePatient, err := connectionEstablished.Prepare("UPDATE patients SET ACTIVE=false WHERE ID=?")
	if err != nil {
		panic(err.Error())
	}
	disablePatient.Exec(idPatient)

	fmt.Printf("El paciente ID (%v) ha sido deshabilitado.\n", idPatient)
}
