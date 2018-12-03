package modelos

import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	"log"
)
//Caixa - CRUD
type Caixa struct {
	ID      int
	Valor   float64
	DataCadastro string
	Login string
}

//CadCaixa 
func CadCaixa(Caixa Caixa) {

	fmt.Println("Cad Caixa")
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	Caixa.DataCadastro = controler.PegarDataAtualString()

	stmt, err := db.Prepare("insert into Caixa(valor,datacadastro,login) values($1,$2,$3)")
	res, err := stmt.Exec(Caixa.Valor, Caixa.DataCadastro,Caixa.Login)
	
	
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
func GetCaixaAll() []Caixa {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,valor,datacadastro,login FROM caixa;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Caixa
	for rows.Next() {
	var co Caixa
	rows.Scan(&co.ID, &co.Valor, &co.DataCadastro,&co.Login)
	ps = append(ps, Caixa{ID: co.ID, Valor: co.Valor, DataCadastro: co.DataCadastro,Login: co.Login})
	}
	defer rows.Close()
	return ps
}

func DeletarCaixa(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from caixa where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	
