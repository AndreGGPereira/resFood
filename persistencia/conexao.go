package persistencia

import (
	 
	"fmt"
	
	"database/sql"
	//banco mysql
	"html/template"
	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template
var err error

func GetConexao()*sql.DB {
	
	fmt.Println("Pegou Conexão")
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/eng?sslmode=disable")
	if err != nil {
	panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println(" Teste conexão You connected to your database.")
	defer db.Close()
	fmt.Println("Entrou no banco")
	return db
}

func AbrirConexao() *sql.DB {

	db, err := sql.Open("mysql", "root:andre110407@/engenharia")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db
}
