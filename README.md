# PROJECT N1: Diagnostic survey for depression.

## Description: 
It is a system for conducting an interview with multivalued responses
(with values ranging from 0 to 3) in order to diagnose symptoms or the disease of depression.

## Project status: 
	Active - On Track.

- Environment Requirements:
	"You have a working Go environment."

## Installing:

Building this system requires some dependencies, this are required and must be installed separately.

	- all versions of Go for Windows are available on https://golang.google.cn/doc/install
	- Download the Golang according to your system architecture and get go1.19.3 version 
	- Execute the downloaded archive file and wherever you want to install this.
	- To check if your device is installed:
		- go to the Command line an run the following command: 
		" go version "
 	- Set the environment variables. Right click on My PC and select Choose the Advanced System Settings
	  from the right side and click on Environment Variables:
		- GOROOT: this variable denote your path instalation C:\Program Files\Go 
		- GOPATH: this variable denote your workspace C:\Usuarios\User_Name\go 
		- GOBIN: this variable denote the bin folder inside of your workspace C:\Usuarios\User_Name\go\bin
	- Edit the following environment variables: 
		- PATH: this variable denote the installation on the bin folder C:\Program Files\Go\bin
	- Open a command line and cd to your proyect directory:
		- $ go mod init sistema
		- $ go get -u github.com/go-sql-driver/mysql 
	###"Steps of conection's system"	
	"Pasos para tener habilitar relaciones en PHPMyAdmin"
	- Iniciar Wampserver, dirigirse a 'MySQL', seleccionar 'my.ini'.
	- Editar 'default-storage'. Debe quedar de la siguiente manera:
		- ;default-storage-engine=MYISAM
		- default-storage-engine=InnoDB
	- Reiniciar Wampserver
	- Crear la base de datos y exportar el archivo sql.
	"Add a list of the minimum steps needed to clone the project and get it running locally."
	

Authorship:
	Mati
	Tsubasa
	-ABS
	GonzaloGo

Licensing:
	CC0 1.0 Universal.
