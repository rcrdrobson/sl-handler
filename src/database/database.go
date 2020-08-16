package database

import (
	"bytes"
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	connection *sql.DB
}

type Function struct {
	Id         int
	Name       string
	Code       string
	Pack       string
	Dockerfile string
}

const (
	sqliteVersion = "sqlite3"
	pathDataBase  = "./database.db"
)

func (d *Database) Connect() {
	var connection, err = sql.Open(sqliteVersion, pathDataBase)
	checkErr(err)
	d.connection = connection
	d.createSchema()
}

func (d *Database) Close() {
	d.connection.Close()
}

func (d *Database) createSchema() {
	switch sqliteVersion {
	case "mysql":
		var qFunctionTable = "CREATE TABLE IF NOT EXISTS function (id INT(10) NOT NULL AUTO_INCREMENT, name TEXT, code TEXT, pack TEXT, dockerfile TEXT, PRIMARY KEY (`id`))"
		statement, _ := d.connection.Prepare(qFunctionTable)
		statement.Exec()
	case "sqlite3":
		var qFunctionTable = "CREATE TABLE IF NOT EXISTS function (id INTEGER PRIMARY KEY, name TEXT, code TEXT, pack TEXT, dockerfile TEXT)"
		statement, _ := d.connection.Prepare(qFunctionTable)
		statement.Exec()
	}
}

func (d *Database) InsertFunction(name string, code string, pack string, dockerfile string) {
	statement, err := d.connection.Prepare("INSERT INTO function (name, code, pack, dockerfile) VALUES (?, ?, ?, ?)")
	checkErr(err)
	_, err = statement.Exec(name, code, pack, dockerfile)
	checkErr(err)
}

func (d *Database) DeleteFunction(name string) bool {
	statement, err := d.connection.Prepare("DELETE FROM function WHERE name=?")
	checkErr(err)

	_, err = statement.Exec(name)
	checkErr(err)

	return err == nil
}

func (d *Database) SelectFunction(name string) string {
	rows, err := d.connection.Query("SELECT * FROM function WHERE name='" + name + "'")
	checkErr(err)
	defer rows.Close()

	if !rows.Next() {
		return ""
	}

	function := Function{}
	err = rows.Scan(&function.Id, &function.Name, &function.Code, &function.Pack, &function.Dockerfile)
	checkErr(err)

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(function)
	return string(buf.Bytes())
}

func (d *Database) SelectAllFunction() string {
	rows, err := d.connection.Query("SELECT * FROM function")
	checkErr(err)
	defer rows.Close()

	var functionList = make([]Function, 0)

	for rows.Next() {
		function := Function{}
		err = rows.Scan(&function.Id, &function.Name, &function.Code, &function.Pack, &function.Dockerfile)
		checkErr(err)
		functionList = append(functionList, function)
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(functionList)
	return string(buf.Bytes())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
