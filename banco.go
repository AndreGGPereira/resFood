package main

// essa classe fará toda a criação das struct no banco

import (
	"database/sql"
	"fmt"
	 //pacote postsql
	_ "github.com/lib/pq"
)

//db e uma referencia do *sql.DB
// returna sql.Result
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println("Teste")
	return result
}


func CriarBancoStruck11() {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	// conecta no banco de dados
	if err != nil {
		panic(err)
	}
	// fecha conexão antes do final do metado
	defer db.Close()

	exec(db, "drop table if exists caixa")
	exec(db, `create table caixa (
		id serial PRIMARY KEY NOT NULL,
		valor real,
		login varchar(80),
		dataCadastro varchar(80))`)

	exec(db, "drop table if exists mesa")
	exec(db, `create table mesa (
		id serial PRIMARY KEY NOT NULL,
		numero integer,
		dataabertura varchar(80),
		datafechamento varchar(80),
		login varchar(80),
		ativa BOOLEAN NOT NULL DEFAULT TRUE,
		valortotal real)`)
	
	exec(db, "drop table if exists permissao")
	exec(db, `create table permissao (
		id serial PRIMARY KEY NOT NULL,
		nome varchar(80),
		descricao varchar(80),
		dataCadastro varchar(80))`)
	
	exec(db, "drop table if exists despesatipo")
	exec(db, `create table despesatipo (
		id serial PRIMARY KEY NOT NULL,
		nome varchar(80),
		descricao varchar(80),
		dataCadastro varchar(80))`)	
		
	exec(db, "drop table if exists sangria")
	exec(db, `create table sangria (
		id serial PRIMARY KEY NOT NULL,
		valor real,
		login varchar(80),
		dataCadastro varchar(80))`)	
	
	exec(db, "drop table if exists entradaTroco")
	exec(db, `create table entradaTroco (
			id serial PRIMARY KEY NOT NULL,
			nome varchar(80),
			valor real,
			login varchar(80),
			dataCadastro varchar(80))`)
	
	
	exec(db, "drop table if exists produtoTipo")
	exec(db, `create table produtoTipo (
		id serial PRIMARY KEY NOT NULL,
		nome varchar(80),
		descricao varchar(80),
		dataCadastro varchar(80))`)

	exec(db, "drop table if exists despesa")
	exec(db, `create table despesa (
		id serial PRIMARY KEY NOT NULL,
		nome varchar(80),
		dataCadastro varchar(80),
		dataPG varchar(80),
		dataVencimento varchar(80),
		valor real,
		login varchar(80),
		pg BOOLEAN NOT NULL DEFAULT FALSE,
		id_despesatipo integer,
		FOREIGN KEY (id_despesatipo) REFERENCES despesatipo (id))`)

	
	exec(db, "drop table if exists produto")
	exec(db, `create table produto (
		id serial PRIMARY KEY NOT NULL,
		nome varchar(80),
		quantidade integer,
		valorUnidade real,
		descricao varchar(80),
		dataCadastro varchar(80),
		id_produtotipo integer,
		FOREIGN KEY (id_produtotipo) REFERENCES produtotipo (id))`)
	
	exec(db, "drop table if exists produtoMesa")
	exec(db, `create table produtoMesa (
		id serial PRIMARY KEY NOT NULL,
		dataCadastro varchar(80),
		quantidade integer,
		valorUnidade real,
		valorTotal real,
		login varchar(80),
		id_mesa integer,
		id_produto integer,
		FOREIGN KEY (id_produto) REFERENCES produto (id),
		FOREIGN KEY (id_mesa) REFERENCES mesa (id))`)


    exec(db, "drop table if exists usuario")
	exec(db, `create table usuario (
		id serial PRIMARY KEY NOT NULL,
		nome varchar(80),
		login varchar(80),
		senha bytea,
		id_permissao integer,
		FOREIGN KEY (id_permissao) REFERENCES permissao (id))`)

	}