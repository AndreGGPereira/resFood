package modelos


import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	
	"log"
)

type EntradaTroco struct {
	ID           int
	Valor 		 float64
	DataCadastro string
	Login string
}


//Cadastro de troca utilizando comit, para prover maior segurança nas transações
func CadEntradaTroco(EntradaTroco EntradaTroco) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	EntradaTroco.DataCadastro = controler.PegarDataAtualString()
	
	rowsCaixa, _ := db.Query("SELECT id,valor,datacadastro FROM Caixa;")

	var ps []Caixa
	var caixa Caixa
	for rowsCaixa.Next() {
		rowsCaixa.Scan(&caixa.ID,&caixa.Valor,&caixa.DataCadastro)
		ps = append(ps, Caixa{ID: caixa.ID,Valor: caixa.Valor,DataCadastro: caixa.DataCadastro})
	}

	caixa.DataCadastro = controler.PegarDataAtualString()	
	caixa.Valor += EntradaTroco.Valor

	tx, _ := db.Begin()

	stmtCaixa, err := tx.Prepare("update Caixa set valor=$1,datacadastro=$2 where id =$3")
	stmtCaixa.Exec(caixa.Valor, caixa.DataCadastro,caixa.ID)

	stmtSan, err := tx.Prepare("insert into entradatroco(valor,datacadastro,login) values($1,$2,$3)")
	stmtSan.Exec(EntradaTroco.Valor, EntradaTroco.DataCadastro,EntradaTroco.Login)
	
	if err != nil {
		fmt.Println("Operação não concluida!!")
		tx.Rollback()
		log.Fatal(err)
	}
	// se tudo ok, passa o comit
	tx.Commit()
	
}
//Lista de entrada de troco
func GetEntradaTrocoAll() []EntradaTroco{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,valor,datacadastro,login FROM entradatroco;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []EntradaTroco
	for rows.Next() {
	var um EntradaTroco
	rows.Scan(&um.ID, &um.Valor, &um.DataCadastro, &um.Login)
	ps = append(ps, EntradaTroco{ID: um.ID, Valor: um.Valor, DataCadastro: um.DataCadastro,Login: um.Login})
	}
	defer rows.Close()
	return ps
}

func DeletarEntradaTroco(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from entradatroco where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	
	