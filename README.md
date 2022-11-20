# PROJECT N1: Diagnostic survey for depression.

## Description 
It is a system for conducting an interview with multivalued responses in order to diagnose symptoms of the depression.


## Project status 
	Active - On Track.


## Installing

Building this system requires some dependencies, this are required and must be installed separately.

### GO
- all versions of Go for Windows are available on https://golang.google.cn/doc/install
- Download the Golang according to your system architecture and get go1.19.3 version 
- Execute the downloaded archive file and wherever you want to install this.
- To check if your device is installed:
	- go to the Command line an run the following command: 
	go version
- Set the environment variables. Right click on My PC and select Choose the Advanced System Settings
from the right side and click on Environment Variables:
	Those two variables on system variables: 
	- GOROOT: this variable denote your path instalation C:\Program Files\Go 	
	- GOBIN: this variable denote the bin folder inside of your workspace C:\Usuarios\User_Name\go\bin
	This one on user variable:
	- GOPATH: this variable denote your workspace C:\Usuarios\User_Name\go 
- Edit the following environment variables: 
	- PATH: this variable denote the installation on the bin folder C:\Program Files\Go\bin
- Open a command line and cd to your proyect directory:

		$ go mod init sistema.
		$ go get -u github.com/go-sql-driver/mysql.
	
	
	
### Wamp
 - all versions of Wamp for Windows are available on https://www.wampserver.com/en/
 - Download the Wamp according to your system architecture and get 3.2.6 version 
	

### Steps of conection's system
"Change the storage mechanism in the database from ISAM to INNODB"
- Start Wampserver
- go to 'MySQL' and select 'my.ini'
- Edit 'default-storage'. It should be as follows	

		-;default-storage-engine=MYISAM.
		-default-storage-engine=InnoDB.
		
- Restart Wampserver
- Crear la base de datos y exportar el archivo sql.
	

## Authorship
	Mati
	Tsubasa
	-ABS
	GonzaloGo

## Licensing
	CC0 1.0 Universal.
