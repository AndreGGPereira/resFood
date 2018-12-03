package modelos

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/andreggpereira/eng/controler"
	"log"
)

type ProdutoTipo struct {
	ID           int
	Nome         string
	Descricao    string
	DataCadastro string
}


//Cadastro Tipo de produto
func CadProdutoTipo(ProdutoTipo ProdutoTipo) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	ProdutoTipo.DataCadastro = controler.PegarDataAtualString()
	
	stmt, err := db.Prepare("insert into ProdutoTipo(nome,descricao,datacadastro) values($1,$2,$3)")
	res, err := stmt.Exec(ProdutoTipo.Nome,ProdutoTipo.Descricao,ProdutoTipo.DataCadastro)
	
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
		panic(err)
			fmt.Println("Operação não concluida!!")
			fmt.Println("Res", res)
	}else{
		fmt.Println("Operação concluida com sucesso!!")
	}
	defer stmt.Close()
}

//Lista de todos os estados
func GetProdutoTipoAll() []ProdutoTipo {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,nome,descricao,datacadastro FROM ProdutoTipo;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []ProdutoTipo

	for rows.Next() {
	var um ProdutoTipo
	rows.Scan(&um.ID, &um.Nome, &um.Descricao, &um.DataCadastro)
	ps = append(ps, ProdutoTipo{ID: um.ID, Nome: um.Nome, Descricao: um.Descricao, DataCadastro: um.DataCadastro})
	}
	defer rows.Close()
	return ps
}
func deletarProdutoTipo(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from produtotipo where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	