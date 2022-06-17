# User account example
A microservices example using GoKit with basic httpAPI, postgreSQL

## Install postgreSQL and pgAdmin 4
1. Download postgreSQL at: 
https://www.postgresql.org/download/
2. Select add-in pgAdmin4 while install postgreSQL or download directly at: 
https://www.pgadmin.org/download/pgadmin-4-windows/


## Create database and table with info:
```
const (
	host   = "localhost"
	port   = 5432
	user   = "anhtuan"
	psswd  = "abc13579"
	dbname = "gokit_example"
)
```

## Ref
Accout service:
```
https://github.com/tensor-programming/go-kit-tutorial/tree/master/account
```
Auth by JWT:
```
https://github.com/eminetto/talk-microservices-gokit
```