package modelos

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	"log"
)

//usuarioTipo - CRUD
type UsuarioTipo struct {
	ID           int
	Nome         string
	Descricao    string
	DataCadastro string
}
type UsuarioTipoList struct{
	AllUsuarioTipo[]UsuarioTipo
}

//CadUsuarioTipo 
func CadUsuarioTipo(UsuarioTipo UsuarioTipo) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	UsuarioTipo.DataCadastro = controler.PegarDataAtualString()

	stmt, err := db.Prepare("insert into UsuarioTipo(nome,descricao,datacadastro) values($1,$2,$2)")
	
	res, err := stmt.Exec(UsuarioTipo.Nome, UsuarioTipo.Descricao, UsuarioTipo.DataCadastro)
	
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
		panic(err)
		fmt.Println("Operação não concluida!!")
	}else{
		fmt.Println("Operação concluida com sucesso!!")
	}
	fmt.Println("Teste res", res)
	defer stmt.Close() // danger!
}

//Lista de todos os estados
func GetUsuarioTipoAll() []UsuarioTipo {

	fmt.Println("Entrou pra lista geral")
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,nome,descricao,datacadastro FROM usuariotipo;")
	
	fmt.Println("Entrou pra lista geral1", rows)
	err = db.Ping()
		if err != nil {
			panic(err)
			
		}
	defer db.Close()

	var ps []UsuarioTipo
	for rows.Next() {
	var ust UsuarioTipo
	rows.Scan(&ust.ID, &ust.Nome, &ust.Descricao, &ust.DataCadastro)
	ps = append(ps, UsuarioTipo{ID: ust.ID, Nome: ust.Nome, Descricao: ust.Descricao, DataCadastro: ust.DataCadastro})
	}
	defer rows.Close()
	return ps
}

func deletarUsuarioTipo(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from usuariotipo where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	