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
	name       string
	fatherName string
	motherName string
	birthday   string
	email      string
	password   string
	active     bool
}

type Patient struct {
	ID          int
	RUN         int
	DV          string
	name        string
	fatherName  string
	motherName  string
	birthday    string
	phone       int
	email       string
	observation string
	active      bool
}

type Test struct {
	ID          int
	name        string
	cutPoint    int
	matchPoint  int
	description string
	active      bool
}

type Question struct {
	ID          int
	ID_t        int
	question    string
	description string
	active      bool
}

type Answer struct {
	ID          int
	ID_q        int
	point       string
	observation string
	active      bool
}

type Poll struct {
	ID   int
	ID_u int
	ID_t int
}

type Session struct {
	ID    int
	ID_po int
	date  string
}

func main() {

	connectionEstablished := conectionBD()

	//Insertions
	//insertUser(connectionEstablished)
	//insertPatient(connectionEstablished)
	//insertTest(connectionEstablished)
	//insertQuestion(connectionEstablished, 1) //Recibe ID Test pero NO confirma si existe ID Test.
	//insertAnswer(connectionEstablished, 1) //Recibe ID Question pero NO confirma si existe ID Question.
	//insertPoll(connectionEstablished, 1, 1) //ID user, ID test.
	//insertSession(connectionEstablished, 1, 1) //ID patient y ID poll.

	//Reciben conexion y una 'ID'
	//disableUser(connectionEstablished, 1)    //No confirma si el usuario existe.
	//disablePatient(connectionEstablished, 1) //No confirma si el usuario existe.
	//disableTest(connectionEstablished, 1)
	//disableQuestion(connectionEstablished, 1)
	//disableAnswer(connectionEstablished, 1)

	//readAllUsers(connectionEstablished)
	//readAllPatients(connectionEstablished)
	//readAllTests(connectionEstablished)
	//readAllQuestions(connectionEstablished)
	readAllAnswers(connectionEstablished)
	//readAllPolls(connectionEstablished)
	//readAllSessions(connectionEstablished)
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

	fmt.Print("Ingresa una descripcion: ")
	o := bufio.NewReader(os.Stdin)
	description, _ := o.ReadString('\n')

	insertTest, err := connectionEstablished.Prepare("INSERT INTO tests (ID_T, NAME_T, CUTPOINT_T, MATCHPOINT_T, DESCRIPTION_T, ACTIVE_T) VALUES (NULL, ?, ?, ?, ?, '1')")
	if err != nil {
		panic(err.Error())
	}
	insertTest.Exec(name, cutPoint, matchPoint, description)

	fmt.Println("Test ingresado.")
}

func insertAnswer(connectionEstablished *sql.DB, ID_question int) {

	var point int
	fmt.Print("Ingresa el puntaje (entre 0 y 3): ")
	fmt.Scanln(&point)

	fmt.Print("Ingresa una observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	insertAnswer, err := connectionEstablished.Prepare("INSERT INTO answers (ID_A, ID_Q, POINT_A, OBSERVATION_A, ACTIVE_A) VALUES (NULL, ?, ?, ?, '1')")
	if err != nil {
		panic(err.Error())
	}
	insertAnswer.Exec(ID_question, point, observation)

	fmt.Println("Respuesta ingresada.")
}

func insertPoll(connectionEstablished *sql.DB, ID_user int, ID_test int) {
	insertPoll, err := connectionEstablished.Prepare("INSERT INTO polls (ID_po, ID_u, ID_t) VALUES (NULL, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertPoll.Exec(ID_user, ID_test)

	fmt.Println("Encuesta ingresada.")
}

func insertSession(connectionEstablished *sql.DB, ID_patient int, ID_poll int) {

	var date string
	fmt.Print("Ingresa una fecha (yyyy-mm-dd): ")
	fmt.Scanln(&date)

	insertSession, err := connectionEstablished.Prepare("INSERT INTO session (ID_p, ID_po, DATE_s) VALUES ( ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertSession.Exec(ID_patient, ID_poll, date)

	fmt.Println("Sesion ingresada.")
}

func readAllUsers(connectionEstablished *sql.DB) {

	fmt.Println("Mostrando a todos los usuarios activos")
	readUser, err := connectionEstablished.Query("SELECT * FROM users WHERE ACTIVE_U=true")
	if err != nil {
		panic(err.Error())
	}

	user := User{} //Variable := Struct
	arrayUser := []User{}

	for readUser.Next() {
		var ID int
		var RUN int
		var DV string
		var name string
		var fatherName string
		var motherName string
		var birthday string
		var email string
		var password string
		var active bool
		err = readUser.Scan(&ID, &RUN, &DV, &name, &fatherName, &motherName, &birthday, &email, &password, &active)
		if err != nil {
			panic(err.Error())
		}

		user.ID = ID
		user.RUN = RUN
		user.DV = DV
		user.name = name
		user.fatherName = fatherName
		user.motherName = motherName
		user.email = email
		user.birthday = birthday
		user.password = password
		user.active = active
		arrayUser = append(arrayUser, user)

	}
	fmt.Println(arrayUser)
}

func readAllPatients(connectionEstablished *sql.DB) {

	fmt.Println("Mostrando a todos los pacientes activos")
	readPatient, err := connectionEstablished.Query("SELECT * FROM patients WHERE ACTIVE_P=true")
	if err != nil {
		panic(err.Error())
	}

	patient := Patient{} //Variable := Struct
	arrayPatient := []Patient{}
	for readPatient.Next() {
		var ID int
		var RUN int
		var DV string
		var name string
		var fatherName string
		var motherName string
		var birthday string
		var phone int
		var email string
		var observation string
		var active bool
		err = readPatient.Scan(&ID, &RUN, &DV, &name, &fatherName, &motherName, &birthday, &phone, &email, &observation, &active)
		if err != nil {
			panic(err.Error())
		}

		patient.ID = ID
		patient.RUN = RUN
		patient.DV = DV
		patient.name = name
		patient.fatherName = fatherName
		patient.motherName = motherName
		patient.birthday = birthday
		patient.phone = phone
		patient.email = email
		patient.observation = observation
		patient.active = active
		arrayPatient = append(arrayPatient, patient)
	}
	fmt.Println(arrayPatient)
}

func readAllTests(connectionEstablished *sql.DB) {

	fmt.Println("Mostrando todos los tests activos")
	readTest, err := connectionEstablished.Query("SELECT * FROM tests WHERE ACTIVE_T=true")
	if err != nil {
		panic(err.Error())
	}

	test := Test{} //Variable := Struct
	arrayTest := []Test{}
	for readTest.Next() {
		var ID int
		var name string
		var cutPoint int
		var matchPoint int
		var description string
		var active bool
		err = readTest.Scan(&ID, &name, &cutPoint, &matchPoint, &description, &active)
		if err != nil {
			panic(err.Error())
		}
		test.ID = ID
		test.name = name
		test.cutPoint = cutPoint
		test.matchPoint = matchPoint
		test.description = description
		test.active = active
		arrayTest = append(arrayTest, test)
	}
	fmt.Println(arrayTest)
}

func readAllQuestions(connectionEstablished *sql.DB) {

	fmt.Println("Mostrando todas las preguntas activas")
	readQuestion, err := connectionEstablished.Query("SELECT * FROM questions WHERE ACTIVE_Q=true")
	if err != nil {
		panic(err.Error())
	}

	questions := Question{} //Variable := Struct
	arrayQuestions := []Question{}
	for readQuestion.Next() {
		var ID int
		var ID_t int
		var question string
		var description string
		var active bool
		err = readQuestion.Scan(&ID, &ID_t, &question, &description, &active)
		if err != nil {
			panic(err.Error())
		}
		questions.ID = ID
		questions.ID_t = ID_t
		questions.question = question
		questions.description = description
		questions.active = active
		arrayQuestions = append(arrayQuestions, questions)
	}
	fmt.Println(arrayQuestions)
}

func readAllAnswers(connectionEstablished *sql.DB) {

	fmt.Println("Mostrando todas las respuestas activas")
	readAnswers, err := connectionEstablished.Query("SELECT * FROM answers WHERE ACTIVE_A=true")
	if err != nil {
		panic(err.Error())
	}

	answer := Answer{} //Variable := Struct
	arrayAnswer := []Answer{}
	for readAnswers.Next() {
		var ID int
		var ID_q int
		var point string
		var observation string
		var active bool
		err = readAnswers.Scan(&ID, &ID_q, &point, &observation, &active)
		if err != nil {
			panic(err.Error())
		}
		answer.ID = ID
		answer.ID_q = ID_q
		answer.point = point
		answer.observation = observation
		answer.active = active
		arrayAnswer = append(arrayAnswer, answer)
	}
	fmt.Println(arrayAnswer)
}

func readAllPolls(connectionEstablished *sql.DB) {
	fmt.Println("Mostrando todas las encuestas")
	readPolls, err := connectionEstablished.Query("SELECT * FROM polls")
	if err != nil {
		panic(err.Error())
	}

	poll := Poll{} //Variable := Struct
	arrayPoll := []Poll{}
	for readPolls.Next() {
		var ID int
		var ID_u int
		var ID_t int
		err = readPolls.Scan(&ID, &ID_u, &ID_t)
		if err != nil {
			panic(err.Error())
		}
		poll.ID = ID
		poll.ID_u = ID_u
		poll.ID_t = ID_t
		arrayPoll = append(arrayPoll, poll)
	}
	fmt.Println(arrayPoll)
}

func readAllSessions(connectionEstablished *sql.DB) {
	fmt.Println("Mostrando todas las sesiones")
	readSessions, err := connectionEstablished.Query("SELECT * FROM session")
	if err != nil {
		panic(err.Error())
	}

	session := Session{} //Variable := Struct
	arraySession := []Session{}
	for readSessions.Next() {
		var ID int
		var ID_po int
		var date string
		err = readSessions.Scan(&ID, &ID_po, &date)
		if err != nil {
			panic(err.Error())
		}
		session.ID = ID
		session.ID_po = ID_po
		session.date = date
		arraySession = append(arraySession, session)
	}
	fmt.Println(arraySession)
}

func disableUser(connectionEstablished *sql.DB, ID int) {

	idUser := ID
	disableUser, err := connectionEstablished.Prepare("UPDATE users SET ACTIVE_u=false WHERE ID_u=?")
	if err != nil {
		panic(err.Error())
	}
	disableUser.Exec(idUser)

	fmt.Printf("El usuario ID (%v) ha sido deshabilitado.\n", idUser)
}

func disablePatient(connectionEstablished *sql.DB, ID int) {

	idPatient := ID
	disablePatient, err := connectionEstablished.Prepare("UPDATE patients SET ACTIVE_p=false WHERE ID_p=?")
	if err != nil {
		panic(err.Error())
	}
	disablePatient.Exec(idPatient)

	fmt.Printf("El paciente ID (%v) ha sido deshabilitado.\n", idPatient)
}

func disableTest(connectionEstablished *sql.DB, ID int) {

	idTest := ID
	disableTest, err := connectionEstablished.Prepare("UPDATE tests SET ACTIVE_t=false WHERE ID_t=?")
	if err != nil {
		panic(err.Error())
	}
	disableTest.Exec(idTest)

	fmt.Printf("El paciente ID (%v) ha sido deshabilitado.\n", idTest)
}

func disableQuestion(connectionEstablished *sql.DB, ID int) {

	idQuestion := ID
	disableQuestion, err := connectionEstablished.Prepare("UPDATE questions SET ACTIVE_q=false WHERE ID_q=?")
	if err != nil {
		panic(err.Error())
	}
	disableQuestion.Exec(idQuestion)

	fmt.Printf("La pregunta ID (%v) ha sido deshabilitado.\n", idQuestion)
}

func disableAnswer(connectionEstablished *sql.DB, ID int) {

	idAnswer := ID
	disableQuestion, err := connectionEstablished.Prepare("UPDATE answers SET ACTIVE_a=false WHERE ID_a=?")
	if err != nil {
		panic(err.Error())
	}
	disableQuestion.Exec(idAnswer)

	fmt.Printf("La respuesta ID (%v) ha sido deshabilitado.\n", idAnswer)
}
