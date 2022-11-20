package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"runtime"

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

	fmt.Println("------------------------------INICIO------------------------------")

	menu(connectionEstablished)

	fmt.Println("--------------------------------FIN-------------------------------")
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
	insertSession, err := connectionEstablished.Prepare("INSERT INTO sessions (ID_p, ID_po, DATE_s) VALUES ( ?, ?, ?)")
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
func readPatient(connectionEstablished *sql.DB, ID_patient int) {

	fmt.Println("Mostrando al paciente")
	readPatient, err := connectionEstablished.Query("SELECT * FROM patients WHERE ID_p= ? ", ID_patient)
	if err != nil {
		panic(err.Error())
	}
	patient := Patient{}
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

func readQuestionsbyAnswer(connectionEstablished *sql.DB, ID_test int) {

	fmt.Println("Mostrando preguntas del test: ")
	readQuestions, err := connectionEstablished.Query("SELECT * FROM questions WHERE ID_t= ?", ID_test)
	if err != nil {
		panic(err.Error())
	}
	questions := Question{} //Variable := Struct
	arrayQuestions := []Question{}
	for readQuestions.Next() {

		var ID int
		var ID_t int
		var question string
		var description string
		var active bool
		err = readQuestions.Scan(&ID, &ID_t, &question, &description, &active)
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

	//fmt.Println("testeando")
	//fmt.Println(arrayQuestions)


	for i := 0; i < len(arrayQuestions); i++ {
		fmt.Println(arrayQuestions[i])
		fmt.Print("------\n")
	}
}

func readQuestionsbyTest(connectionEstablished *sql.DB, ID_test int) {

	fmt.Println("Mostrando preguntas del test: ")
	readQuestions, err := connectionEstablished.Query("SELECT * FROM questions WHERE ID_t= ? ", ID_test)
	if err != nil {
		panic(err.Error())
	}
	questions := Question{} //Variable := Struct
	arrayQuestions := []Question{}
	for readQuestions.Next() {
		var ID int
		var ID_t int
		var question string
		var description string
		var active bool
		err = readQuestions.Scan(&ID, &ID_t, &question, &description, &active)
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
func readAnswersbyQuestion(connectionEstablished *sql.DB, ID_question int) {
	fmt.Println("Mostrando respuestas del test: ")
	readAnswers, err := connectionEstablished.Query("SELECT * FROM answers WHERE ID_q= ? ", ID_question)
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
	readSessions, err := connectionEstablished.Query("SELECT * FROM sessions")
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
func stateUser(connectionEstablished *sql.DB, ID int) {
	var active int
	fmt.Print("Para activar ingrese (1) para desactivar ingrese (0): ")
	fmt.Scanln(&active)
	idUser := ID
	stateUser, err := connectionEstablished.Prepare("UPDATE users SET ACTIVE_u=? WHERE ID_u=?")
	if err != nil {
		panic(err.Error())
	}
	stateUser.Exec(active, idUser)
	if active == 0 {
		fmt.Printf("El usuario ID (%v) ha sido deshabilitado.\n", idUser)
	} else {
		fmt.Printf("El usuario ID (%v) ha sido habilitado.\n", idUser)
	}
}
func statePatient(connectionEstablished *sql.DB, ID int) {
	var active int
	fmt.Print("Para activar ingrese (1) para desactivar ingrese (0): ")
	fmt.Scanln(&active)
	idPatient := ID
	statePatient, err := connectionEstablished.Prepare("UPDATE patients SET ACTIVE_p=? WHERE ID_p=?")
	if err != nil {
		panic(err.Error())
	}
	statePatient.Exec(active, idPatient)
	if active == 0 {
		fmt.Printf("El paciente ID (%v) ha sido deshabilitado.\n", idPatient)
	} else {
		fmt.Printf("El paciente ID (%v) ha sido habilitado.\n", idPatient)
	}
}
func stateTest(connectionEstablished *sql.DB, ID int) {
	var active int
	fmt.Print("Para activar ingrese (1) para desactivar ingrese (0): ")
	fmt.Scanln(&active)
	idTest := ID
	stateTest, err := connectionEstablished.Prepare("UPDATE tests SET ACTIVE_t=? WHERE ID_t=?")
	if err != nil {
		panic(err.Error())
	}
	stateTest.Exec(active, idTest)
	if active == 0 {
		fmt.Printf("El test ID (%v) ha sido deshabilitado.\n", idTest)
	} else {
		fmt.Printf("El test ID (%v) ha sido habilitado.\n", idTest)
	}
}
func stateQuestion(connectionEstablished *sql.DB, ID int) {
	var active int
	fmt.Print("Para activar ingrese (1) para desactivar ingrese (0): ")
	fmt.Scanln(&active)
	idQuestion := ID
	stateQuestion, err := connectionEstablished.Prepare("UPDATE questions SET ACTIVE_q=? WHERE ID_q=?")
	if err != nil {
		panic(err.Error())
	}
	stateQuestion.Exec(active, idQuestion)
	if active == 0 {
		fmt.Printf("La pregunta ID (%v) ha sido deshabilitado.\n", idQuestion)
	} else {
		fmt.Printf("La pregunta ID (%v) ha sido habilitado.\n", idQuestion)
	}
}
func stateAnswer(connectionEstablished *sql.DB, ID int) {
	var active int
	fmt.Print("Para activar ingrese (1) para desactivar ingrese (0): ")
	fmt.Scanln(&active)
	idAnswer := ID
	stateAnswer, err := connectionEstablished.Prepare("UPDATE answers SET ACTIVE_a=? WHERE ID_a=?")
	if err != nil {
		panic(err.Error())
	}
	stateAnswer.Exec(active, idAnswer)
	if active == 0 {
		fmt.Printf("La respuesta ID (%v) ha sido deshabilitado.\n", idAnswer)
	} else {
		fmt.Printf("La respuesta ID (%v) ha sido habilitado.\n", idAnswer)
	}
}

// updates
func updatePatients(connectionEstablished *sql.DB) {

	var ID int
	fmt.Print("Ingrese ID del Paciente: ")
	fmt.Scanln(&ID)

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
	fmt.Print("Ingresa una nueva observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	//idPatients := ID

	updatePatients, err := connectionEstablished.Prepare("UPDATE patients SET (RUN_P = ?, DV_P = ?, NAME_P = ?, FATHERNAME_P = ?, MOTHERNAME_P = ?, PHONE_P = ?, EMAIL_P = ?, BIRTHDAY_P = ?, OBSERVATION_P = ?) WHERE ID_P= ?")
	if err != nil {
		panic(err.Error())
	}
	updatePatients.Exec(RUN, DV, names, fatherName, motherName, phone, email, birthday, observation, ID)
	fmt.Print("Datos actualizados con exito")
}

func updateUsers(connectionEstablished *sql.DB) {

	var ID int
	fmt.Print("Ingrese ID del Usuario: ")
	fmt.Scanln(&ID)

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
	fmt.Print("Ingresa tu contraseña: ")
	fmt.Scanln(&password)

	updateUsers, err := connectionEstablished.Prepare("UPDATE users SET (RUN_U= ?, DV_U, NAME_U= ?, FATHERNAME_U= ?, MOTHERNAME_U= ?, BIRTHDAY_U= ?, EMAIL_U= ?, PASSWORD_U= ?) WHERE ID_U= ?")
	if err != nil {
		panic(err.Error())
	}

	updateUsers.Exec(RUN, DV, names, fatherName, motherName, email, birthday, email, password, ID)
	fmt.Print("Datos actualizados con exito")
}

func updateQuestions(connectionEstablished *sql.DB, ID_test int) {
	var ID_question int
	fmt.Print("Ingresa la ID de la pregunta: ")
	fmt.Scanln(&ID_question)

	fmt.Print("Formula tu pregunta: ")
	q := bufio.NewReader(os.Stdin)
	question, _ := q.ReadString('\n')
	fmt.Print("Escribe su descripcion: ")
	d := bufio.NewReader(os.Stdin)
	description, _ := d.ReadString('\n')

	InsertQuestion, err := connectionEstablished.Prepare("UPDATE questions SET (QUESTION_Q= ?, DESCRIPTION_Q= ?) WHERE (ID_Q = ?); ")
	if err != nil {
		panic(err.Error())
	}

	InsertQuestion.Exec(question, description, ID_question)
	fmt.Println("Pregunta actualizada con exito.")
}

func updateAnswers(connectionEstablished *sql.DB, ID_question int) {
	var idAnswer int
	fmt.Print("Ingresa la ID de la pregunta: ")
	fmt.Scanln(&idAnswer)

	var point int
	fmt.Print("Ingresa el nuevo puntaje (entre 0 y 3): ")
	fmt.Scanln(&point)
	fmt.Print("Ingresa una observacion: ")
	o := bufio.NewReader(os.Stdin)
	observation, _ := o.ReadString('\n')

	insertAnswer, err := connectionEstablished.Prepare("UPDATE answers SET (POINT_A= ?, OBSERVATION_A= ?) WHERE (ID_A= ?)")
	if err != nil {
		panic(err.Error())
	}
	insertAnswer.Exec(point, observation, idAnswer)
	fmt.Println("Respuesta actualizada con exito.")

}

var clear map[string]func()

func init() {
	clear = make(map[string]func())

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func TestBeck() {

}

func menu(connectionEstablished *sql.DB) {

	exitMenu := false

	for exitMenu == false {

		var option int
		exitMenu2 := false

		fmt.Println(" Menu\n Opciones a elegir:\n 1. Crear test o ingresar preguntas segun test.\n 2. Buscar paciente existente.\n 3. Registrar paciente.")
		fmt.Println(" 4. Ver test.\n 5. Responder preguntas.\n 6. Ver ponderacion de paciente.\n 7. Registrar cuestionario.\n 8. Registrar sesion.\n 9. Mas opciones.")
		fmt.Print("Indique su eleccion: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			CallClear()

			var idTest, cantquests, selection int
			fmt.Println("Seleccione alguna de las opciones:\n 1)Crear test\n 2)Agregar Preguntas")
			fmt.Scanln(&selection)
			if selection == 1 {

				insertTest(connectionEstablished)

			} else if selection == 2 {

				fmt.Println("Ingrese ID del test para ingresar las preguntas: ")
				fmt.Scanln(&idTest)

				fmt.Println("Ingrese la cantidad de preguntas deseada para el test: ")
				fmt.Scanln(&cantquests)

				for i := 1; i <= cantquests; i++ {

					fmt.Printf("-----Pregunta numero %d -----\n", i)

					insertQuestion(connectionEstablished, idTest)
				}
			} else {
				fmt.Println("Ingrese una opcion valida.")
			}
			fmt.Println("---------------------------------------")
		case 2:
			CallClear()

			var idPatient int

			fmt.Print("Ingrese ID del paciente a buscar: ")
			fmt.Scanln(&idPatient)

			readPatient(connectionEstablished, idPatient)
			fmt.Println("---------------------------------------")
		case 3:
			CallClear()

			insertPatient(connectionEstablished)
			fmt.Println("---------------------------------------")
		case 4:
			CallClear()

			var idTest int

			fmt.Print("Ingrese ID del test a mostrar: ")
			fmt.Scanln(&idTest)

			readQuestionsbyTest(connectionEstablished, idTest)
			fmt.Println("---------------------------------------")
		case 5:
			CallClear()
			//preguntas

			var idtest int
			var canti int

			fmt.Println("Ingrese la ID del test: ")
			fmt.Scanln(&idtest)
			fmt.Println("-------------")

			readQuestionsbyAnswer(connectionEstablished, idtest)

			fmt.Print("Ingrese la cantidad de respuestas: ")
			fmt.Scanln(&canti)

			fmt.Println("-------------")

			for i := 0; i < canti; i++ {

				var idQuestion int
				fmt.Println("Ingresa la ID de la pregunta: ")
				fmt.Scanln(&idQuestion)

				insertAnswer(connectionEstablished, idQuestion)
				fmt.Println("-------------")

			}

		case 6:

			CallClear()

			var idQuestion int

			fmt.Print("Ingrese ID de la pregunta para ver sus respuestas: ")
			fmt.Scanln(&idQuestion)

			readQuestionsbyTest(connectionEstablished, idQuestion)
			fmt.Println("---------------------------------------")
		case 7:
			CallClear()

			var idUser, idTest int
			fmt.Println("Ingrese ID del usuario para registrar la encuesta: ")
			fmt.Scanln(&idUser)

			fmt.Println("Ingrese ID del test ejecutado: ")
			fmt.Scanln(&idTest)

			insertPoll(connectionEstablished, idUser, idTest)
			fmt.Println("---------------------------------------")
		case 8:
			CallClear()

			var idPatient, idPoll int
			fmt.Println("Ingrese ID del paciente para registrar la sesion: ")
			fmt.Scanln(&idPatient)

			fmt.Println("Ingrese ID de la encuesta ejecutada: ")
			fmt.Scanln(&idPoll)

			insertSession(connectionEstablished, idPatient, idPoll)

			fmt.Println("---------------------------------------")
		case 9:

			for exitMenu2 == false {

				//CallClear()
				var option2 int

				fmt.Println(" Mas Opciones a elegir:\n 1. Crear usuario o ver todos usuarios.\n 2. Ver todos los Pacientes.\n 3. Ver todos los test.")
				fmt.Println(" 4. Ver todas las respuestas.\n 5. Ver todas las preguntas.\n 6. Ver todas las encuestas.\n 7. Ver todas las sesiones.\n 8. Desactivar por ID. \n 9. Salir del sistema.")
				fmt.Print("Indique su eleccion: ")
				fmt.Scanln(&option2)
				switch option2 {
				case 0:
					CallClear()
					exitMenu2 = true
				case 1:
					CallClear()

					var selection int
					fmt.Println("Seleccione alguna de las opciones:\n 1)Crear usuario\n 2)Ver Usuarios")
					fmt.Scanln(&selection)
					if selection == 1 {
						insertUser(connectionEstablished)
					} else if selection == 2 {
						readAllUsers(connectionEstablished)
					} else {
						fmt.Println("Ingrese una opcion valida.")
					}
					fmt.Println("---------------------------------------")
				case 2:
					CallClear()

					readAllPatients(connectionEstablished)

					fmt.Println("---------------------------------------")
				case 3:
					CallClear()

					readAllTests(connectionEstablished)

					fmt.Println("---------------------------------------")
				case 4:
					CallClear()

					readAllQuestions(connectionEstablished)

					fmt.Println("---------------------------------------")
				case 5:
					CallClear()

					readAllAnswers(connectionEstablished)

					fmt.Println("---------------------------------------")
				case 6:
					CallClear()

					readAllPolls(connectionEstablished)

					fmt.Println("---------------------------------------")
				case 7:
					CallClear()

					readAllSessions(connectionEstablished)

					fmt.Println("---------------------------------------")
				case 8:
					CallClear()

					var selection, idstate int
					fmt.Println("Seleccione alguna de las opciones:\n 1)Borrar usuario.\n 2)Borrar Paciente.\n 3)Borrar Test.\n 4)Borrar Pregunta.\n 5)Borrar Respuesta. ")
					fmt.Scanln(&selection)
					if selection == 1 {
						fmt.Print("Indique la Id del usuario a desactivar: ")
						fmt.Scanln(&idstate)

						stateUser(connectionEstablished, idstate)

					} else if selection == 2 {
						fmt.Print("Indique la Id del paciente a desactivar: ")
						fmt.Scanln(&idstate)

						statePatient(connectionEstablished, idstate)

					} else if selection == 3 {
						fmt.Print("Indique la Id del test a desactivar: ")
						fmt.Scanln(&idstate)

						stateTest(connectionEstablished, idstate)

					} else if selection == 4 {
						fmt.Print("Indique la Id de la pregunta a desactivar: ")
						fmt.Scanln(&idstate)

						stateQuestion(connectionEstablished, idstate)

					} else if selection == 5 {
						fmt.Print("Indique la Id del Respuesta a desactivar: ")
						fmt.Scanln(&idstate)

						stateAnswer(connectionEstablished, idstate)

					} else {

						fmt.Println("Ingrese una opcion valida.")

					}
					fmt.Println("---------------------------------------")
				case 9:

					fmt.Println("Saliendo del sistema")
					exitMenu = true
					exitMenu2 = true

				}
			}
		}
	}
}
