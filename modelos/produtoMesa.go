package modelos

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	"log"
)
type ProdutoMesa struct {
	ID             int
	DataCadastro   string
	Quantidade     int
	ValorUnidade   float64
	ValorTotal     float64
	Produto 	   Produto
	Mesa		   Mesa	
	Login 		   string

}

//CadastroProdutoMesa
func CadProdutoMesa(ProdutoMesa ProdutoMesa) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	ProdutoMesa.DataCadastro = controler.PegarDataAtualString()
	stmt, err := db.Prepare("insert into produtomesa(datacadastro,quantidade,valorunidade,valortotal, id_produto, id_mesa,login) values($1,$2,$3,$4,$5,$6,$7)")
	res, err := stmt.Exec(ProdutoMesa.DataCadastro, ProdutoMesa.Quantidade, ProdutoMesa.ValorUnidade, ProdutoMesa.ValorTotal, ProdutoMesa.Produto.ID,ProdutoMesa.Mesa.ID,ProdutoMesa.Login)
	
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
func GetProdutoMesaAll() []ProdutoMesa {

	fmt.Println("ProdutoMesa Estou1")
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,datacadastro,quantidade,valorunidade,valortotal,id_produto,id_mesa,login FROM produtomesa;")

	fmt.Println("Produto Estou2")
	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []ProdutoMesa
	for rows.Next() {
	var co ProdutoMesa

	rows.Scan(&co.ID, &co.DataCadastro, &co.Quantidade,&co.ValorUnidade,&co.ValorTotal, &co.Produto.Nome,&co.Mesa.Numero,&co.Login )
	ps = append(ps, ProdutoMesa{ID: co.ID, DataCadastro: co.DataCadastro, Quantidade: co.Quantidade, ValorUnidade: co.ValorUnidade, ValorTotal: co.ValorTotal, Produto: co.Produto, Mesa: co.Mesa,Login: co.Login})
	}
	defer rows.Close()
	return ps
}
func deletarProdutoMesa(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from produtoMesa where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	
	
