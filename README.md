# PROJECT N1: Diagnostic survey for depression.

## Description: 
	It is a system for conducting an interview with multivalued responses (with values ranging from 0 to 3) in order to diagnose symptoms or the disease of depression.

## Project status: 
	Active - On Track.

- Environment Requirements:
	"You have a working Go environment."

## Installing:

Building this system requires some dependencies, this are required and must be installed separately.

	- all versions of Go for Windows are available on https://golang.google.cn/doc/install
	- Download the Golang according to your system architecture and get go1.19.3 version 
	- Execute the downloaded archive file and wherever you want to install this
		
	"Add a list of the minimum steps needed to clone the project and get it running locally."
	"Pasos para la conexion al sistema:"
 	- Crear las siguientes variables de entorno:
		- GOROOT: esta variable indica su ruta de instalacion C:\Program Files\Go
		- GOPATH: esta variable indica su espacio de trabajo C:\Usuarios\User_Name\go
		- GOBIN: esta variable indica la carpeta bin dentro del espacio de trabajo C:\Usuarios\User_Name\go\bin
	- Editar la siguiente variable de entorno: 
		- PATH: esta variable intida la instalacion en la carpeta bin C:\Program Files\Go\bin
	- Open a command prompt and cd to your proyect directory:
		- $ go mod init sistema
		- $ go get -u github.com/go-sql-driver/mysql 
		
	"Pasos para tener habilitar relaciones en PHPMyAdmin"
	- Iniciar Wampserver, dirigirse a 'MySQL', seleccionar 'my.ini'.
	- Editar 'default-storage'. Debe quedar de la siguiente manera:
		- ;default-storage-engine=MYISAM
		- default-storage-engine=InnoDB
	- Reiniciar Wampserver
	- Crear la base de datos y exportar el archivo sql.

Authorship:
	Mati
	Tsubasa
	-ABS
	GonzaloGo

Licensing:
	CC0 1.0 Universal.
