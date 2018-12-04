package modelos

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	"log"
)

//Cidade - CRUD
type Produto struct {
	ID           int
	Nome         string
	Descricao	 string
	Quantidade   int
	ValorUnidade float64
	DataCadastro string
	ProdutoTipo  ProdutoTipo
}

func CadProduto(Produto Produto) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
			
	Produto.DataCadastro = controler.PegarDataAtualString()
	stmt, err := db.Prepare("insert into produto(nome,Descricao,Quantidade,ValorUnidade,DataCadastro,id_produtoTipo) values($1,$2,$3,$4,$5,$6) RETURNING id")
	//stmt.Exec(estado.Nome, estado.DataCadastro)
	res, err := stmt.Exec(Produto.Nome,Produto.Descricao,Produto.Quantidade,Produto.ValorUnidade,Produto.DataCadastro,Produto.ProdutoTipo.ID)
	
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
func GetProdutoAll() []Produto {

	fmt.Println("Entrou pra lista geral")
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT p.id,p.nome,p.descricao,p.quantidade,p.valorunidade,p.id_produtoTipo, pt.nome FROM produto AS p INNER JOIN produtoTipo AS pt ON p.id_produtotipo = pt.id")
	
	err = db.Ping()
		if err != nil {
			panic(err)
			
		}
	defer db.Close()

	var ps []Produto
	for rows.Next() {
	var ci Produto
	rows.Scan(&ci.ID, &ci.Nome, &ci.Descricao,&ci.Quantidade,&ci.ValorUnidade,&ci.DataCadastro, &ci.ProdutoTipo.ID)
	ps = append(ps, Produto{ID: ci.ID, Nome: ci.Nome,Descricao: ci.Descricao, Quantidade: ci.Quantidade,ValorUnidade: ci.ValorUnidade, DataCadastro: ci.DataCadastro, ProdutoTipo: ci.ProdutoTipo})
	
	}
	defer rows.Close()
	return ps
}

func deletarProduto(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from produto where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}

func GetProdutoIdAll(idProduto int) ([]Produto){

	fmt.Println("Entrou pra lista geral")
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, _ := db.Query("SELECT id, nome,quantidade, valorunidade, id_produtoTipo from produto where id = $1", idProduto)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var ps []Produto
	for rows.Next() {
	var ci Produto
	rows.Scan(&ci.ID, &ci.Nome, &ci.Quantidade,&ci.ValorUnidade,&ci.ProdutoTipo.ID)
	ps = append(ps, Produto{ID: ci.ID, Nome: ci.Nome, Quantidade: ci.Quantidade,ValorUnidade: ci.ValorUnidade, ProdutoTipo: ci.ProdutoTipo})
	
	}
	defer rows.Close()
	return ps
}
	