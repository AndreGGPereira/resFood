package modelos

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	//"github.com/andreggpereira/resFood/controler"
	"log"
	"time"
)

//Usuario capitalize to export from package
type Usuario struct {
	ID        int
	Nome      string
	Login     string
	Senha     []byte
	Permissao Permissao
}

//Session capitalize to export from package
type Session struct {
	Login        string
	LastActivity time.Time
}

func CadUsuario(Usuario Usuario) {
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Id testes ::", Usuario.ID)

	//se o usuario já tiver ID, deverá ser realizada o update
	if(Usuario.ID != 0){
	
	stmt, err := db.Prepare("update Usuario set nome = $1,login = $2, senha = $3, id_permissao = $4 where id =$5")
	stmt.Exec(Usuario.Nome,Usuario.Login, Usuario.Senha,Usuario.Permissao.ID, Usuario.ID )
	
		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			panic(err)
			fmt.Println("Operação não concluida!!")
		//	fmt.Println("Res", res)
		}else{
		fmt.Println("Operação concluida com sucesso!!")
		}

	}else{
	
	stmt, err := db.Prepare("insert into Usuario(nome,login,senha,id_permissao) values($1,$2,$3,$4)")
	res, err := stmt.Exec(Usuario.Nome,Usuario.Login, Usuario.Senha,Usuario.Permissao.ID)

	if err != nil {
		log.Fatal("Cannot run insert statement", err)
		panic(err)
			fmt.Println("Operação não concluida!!")
			fmt.Println("Res", res)
	}else{
		fmt.Println("Operação concluida com sucesso!!")
	}

}
		//defer stmt.Close()
}
//Lista de todos os estados
func GetUsuarioAll() []Usuario {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	rows, err := db.Query("SELECT us.id, us.nome, us.login,us.senha, us.id_permissao, per.nome FROM usuario AS us INNER JOIN permissao AS per ON us.id_permissao = per.id")

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Usuario
	for rows.Next() {
	var um Usuario
	rows.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Permissao.ID,&um.Permissao.Nome)
	ps = append(ps, Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Permissao: um.Permissao})
	}
	return ps
}
func DeletarUsuario(id int) string{

	fmt.Println("Entrou no delete")

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	stmt2, erro := db.Prepare("delete from usuario where id = $1")
	stmt2.Exec(id)

	if(erro != nil){
		return "Operação realizada com sucesso"
	}else{
		return "Não foi possível realizar a operação"
	}
}
func GetUsuarioLogin(login string) (bool,[]Usuario){
	
	fmt.Println("Login", login)
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	resultado ,_ :=  db.Query("select id,nome,login,senha from usuario where login = $1", login)
	//SELECT COUNT(Cliente) AS ClientePaulo FROM Pedidos WHERE Cliente='Paulo'
	//rows, _ := db.Query("select login from usuario where login = $1", login)

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Usuario
	for resultado.Next() {
	var um Usuario
	resultado.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha)
	ps = append(ps, Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha})
	}
	
	if(len(ps) > 0){
		return true , ps
	}else{
		return false, ps
	}
}
func GetUsuarioID(id int) (bool,[]Usuario){
	
	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/resFood?sslmode=disable")
	//SELECT us.id, us.nome, us.login,us.senha, us.id_permissao, per.nome FROM usuario AS us INNER JOIN permissao AS per ON id_permissao = per.id
	resultado ,_ :=  db.Query("SELECT us.id, us.nome, us.login,us.senha, us.id_permissao, per.nome FROM usuario AS us INNER JOIN permissao AS per ON us.id_permissao = per.id AND us.id = $1", id)

	//resultado ,_ :=  db.Query("select id,nome,login,senha,id_permissao from usuario where id = $1", id)

	err = db.Ping()
		if err != nil {
			panic(err)
		}
	defer db.Close()

	var ps []Usuario
	for resultado.Next() {
	var um Usuario
	resultado.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Permissao.ID,&um.Permissao.Nome)
	ps = append(ps, Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Permissao: um.Permissao})
	}
	
	if(len(ps) > 0){
		return true , ps
	}else{
		return false, ps
	}
}
