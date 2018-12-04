package modelos

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/andreggpereira/resFood/controler"
	"log"
)
type Mesa struct {
	ID           int
	Numero       int
	ValorTotal	 float64
	DataAbertura string
	DataFechamento string
	Ativa bool
	Login string
}

func CadMesa(Mesa Mesa) {
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	
	Mesa.DataAbertura = controler.PegarDataAtualString()
	
	stmt, err := db.Prepare("insert into Mesa(numero,valortotal,dataabertura, datafechamento,ativa,login) values($1,$2,$3,$4,$5,$6)")
	res, err := stmt.Exec(Mesa.Numero,Mesa.ValorTotal,Mesa.DataAbertura, Mesa.DataFechamento, Mesa.Ativa,Mesa.Login)
	
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
func GetMesaAll() []Mesa {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id,numero,valortotal,dataAbertura,datafechamento,ativa,login FROM Mesa where ativa=true;")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Mesa

	for rows.Next() {
	var um Mesa
	rows.Scan(&um.ID, &um.Numero, &um.ValorTotal, &um.DataAbertura, &um.DataFechamento, &um.Ativa,&um.Login)
	ps = append(ps, Mesa{ID: um.ID, Numero: um.Numero, ValorTotal: um.ValorTotal, DataAbertura: um.DataAbertura, DataFechamento: um.DataFechamento, Ativa: um.Ativa, Login: um.Login})
	}
	defer rows.Close()
	return ps
}

func GetMesaNumero(numero int) bool {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT id FROM Mesa where ativa=true AND numero=$1", numero)
	if err != nil {
			panic(err)
		}
	defer db.Close()

	if(rows.Next() == false){
		return false
	
	}else{
		return true
	}
}

func deletarMesa(id int) string{

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from mesa where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}

func GetMesaID(id int) (bool,[]Mesa){
	
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	resultado ,_ :=  db.Query("SELECT id,numero,valortotal,dataabertura,datafechamento,ativa,login FROM mesa where id = $1", id)

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Mesa
	for resultado.Next() {
	var um Mesa
	resultado.Scan(&um.ID, &um.Numero, &um.ValorTotal, &um.DataAbertura, &um.DataFechamento,&um.Ativa,&um.Login)
	ps = append(ps, Mesa{ID: um.ID, Numero: um.Numero, ValorTotal: um.ValorTotal, DataAbertura: um.DataAbertura, DataFechamento: um.DataFechamento, Ativa: um.Ativa, Login: um.Login})
	}
	
	if(len(ps) > 0){
		return true , ps
	}else{
		return false, ps
	}
}


func FecharMesa(Mesa Mesa){

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()
	
	rowsCaixa, _ := db.Query("SELECT id,valor,datacadastro,login FROM Caixa")
	
	var ps []Caixa
	var ca Caixa
	for rowsCaixa.Next() {
	rowsCaixa.Scan(&ca.ID, &ca.Valor, &ca.DataCadastro,&ca.Login)
	ps = append(ps, Caixa{ID: ca.ID, Valor: ca.Valor, DataCadastro: ca.DataCadastro,Login: ca.Login})
	}

	ca.Valor += Mesa.ValorTotal
	
	tx, _ := db.Begin()
	
	Mesa.DataFechamento = controler.PegarDataAtualString()
	stmtMesa, err := tx.Prepare("update  mesa set ativa=$1 ,datafechamento=$2 ,login=$3 where id =$4")
	stmtMesa.Exec(Mesa.Ativa, Mesa.DataFechamento,Mesa.Login,Mesa.ID)

	fmt.Println("Teste Mesa",Mesa.Ativa, Mesa.DataFechamento,Mesa.Login,Mesa.ID)
	

	stmtCaixa, err := tx.Prepare("update caixa set valor=$1 where id =$2")
	stmtCaixa.Exec(ca.Valor,ca.ID)

	if err != nil {
		fmt.Println("Operação não concluida!!")
		tx.Rollback()
		log.Fatal(err)
	}
	// se tudo ok, passa o comit
	tx.Commit()
}




