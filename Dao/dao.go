package Dao

import (
	"database/sql"
)

const DB_DRIVER string = "mysql"
const DB_USER_NAME string = "root"
const DB_USER_PASS string = "SQLserver5728@"
const DB_ADDRESS string = "localhost"
const DB_PORT string = "3306"
const DB_NAME string = "kms_gaosi"

func Connect() *sql.DB{
	var db *sql.DB
	var err error
	db, err = sql.Open(DB_DRIVER, DB_USER_NAME+":"+DB_USER_PASS+"@"+"tcp("+DB_ADDRESS+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
		panic(err)
	}
	return db
}