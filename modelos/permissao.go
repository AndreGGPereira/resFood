package modelos

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/andreggpereira/eng/controler"
	"log"
)
//banco ok
//Permissao tipos de permissão do Usuario
type Permissao struct {
	ID           int
	Nome         string
	Descricao    string
	DataCadastro string

}
//CadastroPermissao
func CadPermissao(Permissao Permissao) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	Permissao.DataCadastro = controler.PegarDataAtualString()
	
	stmt, err := db.Prepare("insert into Permissao(nome,descricao,datacadastro) values($1,$2,$3)")
	res, err := stmt.Exec(Permissao.Nome,Permissao.Descricao, Permissao.DataCadastro)
	
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
		panic(err)
			fmt.Println("Operação não concluida!!")
	}else{
		fmt.Println("Operação concluida com sucesso!!")
	}
	fmt.Println("Teste res", res)

	defer stmt.Close()
	
}
//Lista de todos os estados
func GetPermissaoALL() []Permissao {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,nome,descricao,datacadastro FROM Permissao;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Permissao
	for rows.Next() {
	var um Permissao
	rows.Scan(&um.ID, &um.Nome, &um.Descricao, &um.DataCadastro)
	ps = append(ps, Permissao{ID: um.ID, Nome: um.Nome, Descricao: um.Descricao, DataCadastro: um.DataCadastro})
	}
	defer rows.Close()
	return ps
}
func deletarPermissao(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from permissao where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	
	


