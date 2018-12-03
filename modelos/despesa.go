package modelos

import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq"
	"github.com/andreggpereira/eng/controler"
	"log"
)
//Despesa - CRUD
type Despesa struct {
	ID             int
	Nome           string
	DataCadastro   string
	Descricao      string
	DespesaTipo    DespesaTipo
	Valor          float64
	DataDespesa    string
	DataPG         string
	DataVencimento string
	Pago           bool
	Login string
}

//CadastroEstado 
func CadDespesa(Despesa Despesa) {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	Despesa.DataCadastro = controler.PegarDataAtualString()
	stmt, err := db.Prepare("insert into Despesa(nome,datacadastro,descricao,id_despesaTipo,valor, datadespesa,datapg, datavencimento,pago,login) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)")
	res, err := stmt.Exec(Despesa.Nome, Despesa.DataCadastro, Despesa.Descricao, Despesa.DespesaTipo.ID, Despesa.Valor, Despesa.DataDespesa,Despesa.DataPG, Despesa.DataVencimento,Despesa.Pago,Despesa.Login)
	
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

//Lista de todos das despesas
func GetDespesaAll() []Despesa {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,nome,datacadastro,descricao,id_despesaTipo,valor, datadespesa,datapg, datavencimento,pago,login FROM Despesa;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Despesa
	for rows.Next() {
	var co Despesa

	rows.Scan(&co.ID, &co.Nome, &co.DataCadastro,&co.Descricao, &co.DespesaTipo.Nome, &co.Valor,&co.DataDespesa, &co.DataPG, &co.DataVencimento, &co.Pago,&co.Login)
	ps = append(ps, Despesa{ID: co.ID, Nome: co.Nome, DataCadastro: co.DataCadastro, Descricao: co.Descricao, DespesaTipo: co.DespesaTipo, Valor: co.Valor, DataDespesa: co.DataDespesa, DataPG: co.DataPG, DataVencimento: co.DataVencimento, Pago: co.Pago,Login: co.Login})
	}
	defer rows.Close()
	return ps
}

func deletarDespesa(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from despesa where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
	
	
