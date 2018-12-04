#ResFood - Abaixo temos alguns trexos de código

#Importação de pacotes
	*crypto/hmac"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"time"
	"strconv"
	"github.com/andreggpereira/resFood/modelos"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"



#func init() código que será iniciado ao executar o software 
	//Faz com que o template possa acessar todas as paginas da pasta templates com a teminação .gohtml
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	// metado faz a chamada para criação do banco de dados
	//	CriarBancoStruck()


#Metado main
	// Temos as rotas e a metado que casa rota deve acessar ao ser acessada
	http.HandleFunc("/usuario", usuario)
	http.HandleFunc("/login", login)
	//Metado que inicia o servidor Web 
	http.ListenAndServe(":8080", nil)

#Criação de uma Struct
type Usuario struct {
	ID        int
	Nome      string
	Login     string
	Senha     []byte
	Permissao Permissao
}
#CRUD
Função de Cadastro

	*stmt, err := db.Prepare("insert into Usuario(nome,login,senha,id_permissao) values($1,$2,$3,$4)")
	*res, err := stmt.Exec(Usuario.Nome,Usuario.Login, Usuario.Senha,Usuario.Permissao.ID)

Função de Update

	*stmt, err := db.Prepare("update Usuario set nome = $1,login = $2, senha = $3, id_permissao = $4 where id =$5")
	*stmt.Exec(Usuario.Nome,Usuario.Login, Usuario.Senha,Usuario.Permissao.ID, Usuario.ID )

Função de Delete

	*stmt2, erro := db.Prepare("delete from usuario where id = $1")
	*stmt2.Exec(id)


Função de Ler
	
	*rows, err := db.Query("SELECT * FROM usuario")
	// popular o resultado em um array da strucy
	*var usuarios []Usuario
	*for rows.Next() {
	*var um Usuario
	*rows.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Permissao.ID,&um.Permissao.Nome)
	*usuarios = append(ps, Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Permissao: um.Permissao})
	*}


#HASH - Abaixo temos alguns trexos de código
	//Devemos importar os pacotes abaixo
	*"crypto/hmac"
	*"crypto/sha256"
	*"golang.org/x/crypto/bcrypt"
	//Essa função gera um Hash da senha informada
	*bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	//Já esse metado vamos fazer a comparação do Hash da senha que foi inserida com a que esta cadastrada
	*err := bcrypt.CompareHashAndPassword(u.Senha, []byte(p))

#Controle de Acesso - Criar a sessao
	*	sID, _ := uuid.NewV4()
	*	c := &http.Cookie{
	*		Name:  "session",
	*		Value: sID.String(),
	*	}
	*	c.MaxAge = sessionLength
	*	http.SetCookie(w, c)
	*	dbSessions[c.Value] = session{un, time.Now()}

#Controle de Acesso - Verificando de usuário esta logado para ter acesso

	*c, err := req.Cookie("session")
	*if err != nil {
	*	return false
	*}
	*s, ok := dbSessions[c.Value]
	*if ok {
	*	s.lastActivity = time.Now()
	*	dbSessions[c.Value] = s
	*}
	*c.MaxAge = sessionLength
	*http.SetCookie(w, c)
	*return ok


#Integração com HTML
	//Pegamos informações vindas do formulário
	* nome := req.FormValue("nome")
	
	//Enviando lista de usuarios para pagina HTML
	*type TodoPageData struct {
	*Usuarios []modelos.Usuario
	*}

	*p1 := TodoPageData{Usuarios : modelos.GetUsuarioAll()}
	//Chama o template passando a pagina e a struct
	*tpl.ExecuteTemplate(w, "usuario.gohtml", p1)
	
	//Listagem dos Usuarios na pagina HTML

	*<table>
       *{{range .Usuarios}}
       *<tr>
       *<td>{{.ID}}  </td>
       *<td>{{.Nome}}  </td>
       *<td>{{.Login}}  </td>
       *</tr>
       *{{end}}
	*</table>

#Realizando Commit
	*tx, _ := db.Begin()

	*stmtCaixa, err := tx.Prepare("update Caixa set valor=$1,datacadastro=$2 where id =$3")
	*stmtCaixa.Exec(caixa.Valor, caixa.DataCadastro,caixa.ID)

	*stmtSan, err := tx.Prepare("insert into sangria(valor,datacadastro,login) values($1,$2,$3)")
	*stmtSan.Exec(Sangria.Valor, Sangria.DataCadastro,Sangria.Login)
	
	*if err != nil {
	*	fmt.Println("Operação não concluida!!")
	*	tx.Rollback()
	*	log.Fatal(err)
	*}
	// se tudo ok, passa o comit
	*tx.Commit()

#Banco de Dados
	//Criar Banco
	*exec(db, "create database if not exists resFood")
	*exec(db, "use resFood")

	//Tabela Usuario
   	*exec(db, "drop table if exists usuario")
	*exec(db, `create table usuario (
	*	id serial PRIMARY KEY NOT NULL,
	*	nome varchar(80),
	*	login varchar(80),
	*	senha bytea,
	*	id_permissao integer,
	*	FOREIGN KEY (id_permissao) REFERENCES permissao (id))`)

	













