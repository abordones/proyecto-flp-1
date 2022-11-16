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
// UPDATE

func main() {
	connectionEstablished := conectionBD()
	//-----------------------------INSERTION-----------------------------
	//insertUser(connectionEstablished)
	//insertPatient(connectionEstablished)
	//insertTest(connectionEstablished)
	//insertQuestion(connectionEstablished, 1) //Recibe ID Test.
	//insertAnswer(connectionEstablished, 1) //Recibe ID Question.
	//insertPoll(connectionEstablished, 1, 1) //Recibe ID user, ID test.
	//insertSession(connectionEstablished, 1, 1) //Recibe ID patient y ID poll.
	//-------------------------------STATE-------------------------------
	//stateUser(connectionEstablished, 1)
	//statePatient(connectionEstablished, 1)
	//stateTest(connectionEstablished, 1)
	//stateQuestion(connectionEstablished, 1)
	//stateAnswer(connectionEstablished, 1)
	//-------------------------------READ-------------------------------
	//readAllUsers(connectionEstablished)
	//readAllPatients(connectionEstablished)
	//readAllTests(connectionEstablished)
	//readAllQuestions(connectionEstablished)
	//readAllAnswers(connectionEstablished)
	//readAllPolls(connectionEstablished)
	//readAllSessions(connectionEstablished)
	fmt.Println("------------------------------INICIO------------------------------")
	//1.A. Crear test.
	//insertTest(connectionEstablished) //Como conocer el ID del test generado?
	//2.A. Agregarle preguntas.
	//insertQuestion(connectionEstablished, 2) //Recibe ID Test. //Como conocer el ID del test generado?
	//insertQuestion(connectionEstablished, 2) //Recibe ID Test. //Como conocer el ID del test generado?
	//insertQuestion(connectionEstablished, 2) //Recibe ID Test. //Como conocer el ID del test generado?
	//insertQuestion(connectionEstablished, 2) //Recibe ID Test. //Como conocer el ID del test generado?
	//1.B. Ingresar Usuario.
	//insertUser(connectionEstablished) //Como conocer su ID?
	//2.B. Imprimir TEST.
	//Necesito ID Test e imprimir desde Questions (WHERE ID_T='?').
	//3.B. Agregar respuestas.
	//Imprimir question segun ID?
	//insertAnswer(connectionEstablished, 1) //Recibe ID Question. //Como conocer el ID question generado?
	//insertAnswer(connectionEstablished, 2) //Recibe ID Question. //Como conocer el ID question generado?
	//insertAnswer(connectionEstablished, 3) //Recibe ID Question. //Como conocer el ID question generado?
	//insertAnswer(connectionEstablished, 4) //Recibe ID Question. //Como conocer el ID question generado?
	//4.B. Actualizar matchPoint de TEST.
	//Comprararlo con cutPointy dar resultado -> Deperesion/Normal.
	//5.B. Ingresar POll.
	//insertPoll(connectionEstablished, 1, 2) //Recibe ID user, ID test. //Se requiere esa info.
	//6.B. Ingresar SESSION. 
	//insertSession(connectionEstablished, 1, 2) //Recibe ID patient y ID poll. //Se requiere esa info.

	//INNOBD
	readAllUsers(connectionEstablished)
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

func menu(connectionEstablished *sql.DB) {

	var option int
	fmt.Print(" ")
	fmt.Print("Ingrese una opcion a elegir:\n 1.Crear test o ingresar preguntas segun test\n 2. Buscar paciente existente\n 4. Registrar paciente\n")
	fmt.Print(" 5. Ver test\n 6. Responder preguntas\n 7. Ver ponderacion de paciente\n 8. Ver informacion paciente\n 9. Mas opciones\n ")
	fmt.Scanln(&option)

	//1.1 usuario crea test//
	//1.2 usuario ingresa preguntas//
	//2 usuario busca paciente si es que existe
	//3 usuario registrara a paciente
	//4 llamar print test
	//5 paciente solo respondera preguntas
	//6 se printearan los puntajes segun paciente
	//7 print info paciente
	//8 mas opciones
	//9 CRUD

	switch option {
	case 1:
		var idtest, cantquests, selection int
		fmt.Println("Seleccione alguna de las opciones:\n 1)Crear test\n 2)Agregar Preguntas")
		fmt.Scanln(&selection)
		if selection == 1 {

			insertTest(connectionEstablished)
			//primary := "PRIMARY"
			//r/eadIDTest, err := connectionEstablished.Prepare("SHOW KEYS FROM TESTS WHERE ID_t = ? ")
			//if err != nil {
			//	panic(err.Error())
			//}
			//readIDTest.Exec(primary)
		} else if selection == 2 {

			fmt.Println("Ingrese ID del test para ingresar las preguntas: ")
			fmt.Scanln(&idtest)

			fmt.Println("Ingrese la cantidad de preguntas deseada para el test: ")
			fmt.Scanln(&cantquests)

			for i := 1; i <= cantquests; i++ {

				fmt.Printf("-----Pregunta numero %d -----\n", i)

				insertQuestion(connectionEstablished, idtest)
			}
		} else {
			fmt.Println("Ingrese una opcion valida")
		}

	case 2:

		var idPatient int
		fmt.Println("Ingrese ID del paciente a buscar: ")
		fmt.Scanln(&idPatient)

		readPatient(connectionEstablished, idPatient)

	case 3:

		insertPatient(connectionEstablished)

	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	}

	