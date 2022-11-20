# PROJECT N1: Diagnostic survey for depression

## Description 
It is a system for conducting an interview with multivalued responses in order to diagnose symptoms of depression.


## Project status 
	Active - On Track


## Installing

Building this system requires some dependencies, this are required and must be installed separately.

### GO
- All versions of Go for Windows are available on https://golang.google.cn/doc/install
- Download the Golang according to your system architecture and get go1.19.3 version.
- Execute the downloaded archive file and wherever you want to install this.
- To check if your device is installed, go to the Command line an run the following command:

		go version	
	
- Set the environment variables. Right click on My PC and select Choose the Advanced System Settings.
- From the right side and click on Environment Variables:

1. Those two variables on system variables: 

		GOROOT: this variable denote your path instalation C:\Program Files\Go 	
		GOBIN: this variable denote the bin folder inside of your workspace C:\Usuarios\User_Name\go\bin
	
2. This one on user variable:

		GOPATH: this variable denote your workspace C:\Usuarios\User_Name\go 
	
3. Edit the following environment variables: 

		PATH: this variable denote the installation on the bin folder C:\Program Files\Go\bin

4. Open a command line and cd to your proyect directory:

		$ go mod init sistema.
		$ go get -u github.com/go-sql-driver/mysql.
	
	
	
### Wamp
 - All versions of Wamp for Windows are available on https://www.wampserver.com/en/
 - Download the Wamp according to your system architecture and get 3.2.6 version.
	

##### Change the storage mechanism from ISAM to INNODB
1. Start Wampserver (once started its icon will appear in hidden icons).
2. Go to `MySQL` and select `my.ini`
3. Edit `default-storage`. It should be as follows.	

		;default-storage-engine=MYISAM.
		default-storage-engine=InnoDB.
		
4. Restart Wampserver


### PHP My Admin
1. Start Wampserver (once started its icon will appear in hidden icons).
2. Select `phpMyAdmin`
3. Log in using the following default account.

		Username: root
		Password: 
		
4. Go to create a data base and name it `bd-flp-1`
5. Import the SQL file `BD.sql`
	

## Authorship
	Mati
	Tsubasa
	-ABS
	GonzaloGo

## Licensing
	CC0 1.0 Universal.
