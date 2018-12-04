package modelos


import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	
	"log"
)

type Sangria struct {
	ID           int
	Valor 		 float64
	DataCadastro string
	Login string
}
//Cadastro de retirada de dinheiro, e subtração no caixa com commit
func CadSangria(Sangria Sangria) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	Sangria.DataCadastro = controler.PegarDataAtualString()
	
	rowsCaixa, _ := db.Query("SELECT id,valor,datacadastro FROM Caixa;")

	var ps []Caixa
	var caixa Caixa
	for rowsCaixa.Next() {
		rowsCaixa.Scan(&caixa.ID,&caixa.Valor,&caixa.DataCadastro)
		ps = append(ps, Caixa{ID: caixa.ID,Valor: caixa.Valor,DataCadastro: caixa.DataCadastro})
	}


	if(caixa.Valor < Sangria.Valor ){
		fmt.Println("Valor não disponivel em caixa!!")
		return
	}

	caixa.DataCadastro = controler.PegarDataAtualString()	
	caixa.Valor -= Sangria.Valor

	tx, _ := db.Begin()

	stmtCaixa, err := tx.Prepare("update Caixa set valor=$1,datacadastro=$2 where id =$3")
	stmtCaixa.Exec(caixa.Valor, caixa.DataCadastro,caixa.ID)

	stmtSan, err := tx.Prepare("insert into sangria(valor,datacadastro,login) values($1,$2,$3)")
	stmtSan.Exec(Sangria.Valor, Sangria.DataCadastro,Sangria.Login)
	
	if err != nil {
		fmt.Println("Operação não concluida!!")
		tx.Rollback()
		log.Fatal(err)
	}
	// se tudo ok, passa o comit
	tx.Commit()
	
}
//Lista de todos os estados
func GetSangriaAll() []Sangria{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,valor,datacadastro,login FROM Sangria;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Sangria
	for rows.Next() {
	var um Sangria
	rows.Scan(&um.ID, &um.Valor, &um.DataCadastro, &um.Login)
	ps = append(ps, Sangria{ID: um.ID, Valor: um.Valor, DataCadastro: um.DataCadastro,Login: um.Login})
	}
	defer rows.Close()
	return ps
}
func DeletarSangria(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from sangria where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	
	