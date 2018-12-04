package modelos

import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	"log"
)
//Servi como classe assistente para o cadastro da DespesaTipo
type DespesaTipo struct {
	ID           int
	Nome         string
	Descricao    string
	DataCadastro string
}

//CadastroEstado 
func CadDespesaTipo(DespesaTipo DespesaTipo) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	DespesaTipo.DataCadastro = controler.PegarDataAtualString()
	
	stmt, err := db.Prepare("insert into DespesaTipo(nome,descricao,datacadastro) values($1,$2,$3) RETURNING id")
	//stmt.Exec(estado.Nome, estado.DataCadastro)
	res, err := stmt.Exec(DespesaTipo.Nome,DespesaTipo.Descricao ,DespesaTipo.DataCadastro)
	
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
func GetDespesaTipoAll() []DespesaTipo {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,nome,datacadastro FROM despesatipo;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []DespesaTipo
	for rows.Next() {
	var de DespesaTipo
	rows.Scan(&de.ID, &de.Nome, &de.DataCadastro)
	ps = append(ps, DespesaTipo{ID: de.ID, Nome: de.Nome, Descricao: de.Descricao, DataCadastro: de.DataCadastro})
	}
	defer rows.Close()
	return ps
}
func deletarDespesaTipo(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from despesatipo where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	

