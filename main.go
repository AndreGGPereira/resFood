package main

import (
	"crypto/hmac"
	"crypto/sha256"
	//"database/sql"
	"fmt"
	"html/template"
	//"io"
	"log"
	"net/http"
	"time"
	"strconv"
	"github.com/andreggpereira/resFood/modelos"
	//"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	//pacote postsql
	_ "github.com/lib/pq"
	//"reflect"
)
//para importar pacotes = go get "nome do pacote" ex: go get github.com/julienschmidt/httprouter
//o pacote será baixado e instalado no seu GO, e em todos os seus projetos basta colocar o import
//



// struct da session
type session struct {
	un           string
	lastActivity time.Time
}

type browseModel struct {
	Bucket     string
	Folder     string
}

type Todo struct {
	Title string
	Done bool
	Nome string
}

//para usuar arquivos dentro de outro pacote, basta importar e fazer a chamada do pacote
//"github.com/andreggpereira/resFood/modelos", para chamar qualquer arquivo GO dentro desse pacote do GIT
//basta fazer a chamada = "modelos.Caixa"


type TodoPageData struct {
	Titulo string
	Caixa []modelos.Caixa
	Despesa []modelos.Despesa
	DespesaTipo []modelos.DespesaTipo
	EntradaTroco []modelos.EntradaTroco
	Mesa []modelos.Mesa
	Permissao []modelos.Permissao
	Produto []modelos.Produto
	ProdutoMesa []modelos.ProdutoMesa
	ProdutoTipo []modelos.ProdutoTipo
	Sangria []modelos.Sangria
	Usuarios []modelos.Usuario
	UsuarioTipo []modelos.UsuarioTipo
	Usuario modelos.Usuario
}

type TodoPermissao struct {
	Titulo string
}
//variavel com ponteiro do template
var tpl *template.Template
var dbUsers = map[string]modelos.Usuario{} // user ID, user
var dbSessions = map[string]session{}      // session ID, session
var dbSessionsCleaned time.Time

//PassarDados string
const sessionLength int = 60


//metado executado ao iniciar o aplicativos
func init() {
	//tpl, _ = template.ParseGlob("templates/*.gohtml")
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	dbSessionsCleaned = time.Now()

	// metado faz a chamada para ciação da classe
	//	CriarBancoStruck()
}

func main() {

	//metados que seram chamados ao acessar as rotas
	http.HandleFunc("/", index)
	http.HandleFunc("/caixa", caixa)
	http.HandleFunc("/delusuario", delusuario)
	http.HandleFunc("/altusuario", altusuario)
	http.HandleFunc("/despesa", despesa)
	http.HandleFunc("/despesatipo", despesatipo)
	http.HandleFunc("/entradatroco", entradatroco)
	http.HandleFunc("/mesa", mesa)
	http.HandleFunc("/permissao", permissao)
	http.HandleFunc("/produto", produto)
	http.HandleFunc("/produtomesa", produtomesa)
	http.HandleFunc("/produtotipo", produtotipo)
	http.HandleFunc("/sangria", sangria)
	http.HandleFunc("/usuario", usuario)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
// Servidor funcionará na porta 8080


}

//func index(w http.ResponseWriter, req *http.Request) {
//	u := getUser(w, req)
//	showSessions() // for demonstration purposes
//	tpl.ExecuteTemplate(w, "index.gohtml", u)
//}

//Executa Pagina para fazer login
func index(w http.ResponseWriter, req *http.Request) {
		u := getUser(w, req)
		showSessions() // for demonstration purposes
		tpl.ExecuteTemplate(w, "index.gohtml", u)

}
func caixa(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Caixassaa")

	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var Caixa modelos.Caixa

	if req.Method == http.MethodPost {
		fmt.Println("Caixassaa2",req.FormValue("valor"))
	fltStr,err := strconv.ParseFloat(req.FormValue("valor"),64)

		if(err != nil){
			http.Error(w, "Não possível realizar ", http.StatusForbidden)
			return
	}

	
		//pegar usuario da sessao
	for k, v := range dbSessions {
		Caixa.Login= v.un
		fmt.Println("Teste Usuario",k, v.un)
	}
	
		Caixa.Valor = fltStr
		modelos.CadCaixa(Caixa)		
	}
	p1 := TodoPageData{Caixa : modelos.GetCaixaAll()}
	tpl.ExecuteTemplate(w, "caixa.gohtml", p1)
}
func despesa(w http.ResponseWriter, req *http.Request) {
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var Despesa modelos.Despesa
	if req.Method == http.MethodPost {

		for k, v := range dbSessions {
			Despesa.Login= v.un
			fmt.Println("Teste Usuario",k, v.un)
		}
		Despesa.Nome = req.FormValue("nome")
		Despesa.Descricao = req.FormValue("descricao")
		Despesa.DataDespesa = req.FormValue("datadespesa")
		Despesa.DataPG = req.FormValue("datapg")
		Despesa.DataVencimento = req.FormValue("datavencimento")
		
		if(req.FormValue("pago")=="true"){
			Despesa.Pago = true
		}else{
			Despesa.Pago = false
		}
		
		if s, err := strconv.ParseInt(req.FormValue("despesaTipo"), 10, 64); err == nil {
			Despesa.DespesaTipo.ID = int(s)
		}
		
		fltStr,err := strconv.ParseFloat(req.FormValue("valor"),64)
		if(err != nil){
			http.Error(w, "Não possível realizar ", http.StatusForbidden)
			return
		}
		
		Despesa.Valor = fltStr
		
		modelos.CadDespesa(Despesa)		
	}

	//showSessions() 
	p1 := TodoPageData{Despesa : modelos.GetDespesaAll()}
		tpl.ExecuteTemplate(w, "despesa.gohtml", p1)
}
func despesatipo(w http.ResponseWriter, req *http.Request) {
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	
	var DespesaTipo modelos.DespesaTipo
	if req.Method == http.MethodPost {
		DespesaTipo.Nome = req.FormValue("nome")
		DespesaTipo.Descricao = req.FormValue("descricao")
		modelos.CadDespesaTipo(DespesaTipo)		
	}
	//showSessions() 
	p1 := TodoPageData{DespesaTipo : modelos.GetDespesaTipoAll()}

	tpl.ExecuteTemplate(w, "despesatipo.gohtml", p1)
}
func entradatroco(w http.ResponseWriter, req *http.Request) {
	var EntradaTroco modelos.EntradaTroco
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//pegar usuario da sessao
	for k, v := range dbSessions {
		EntradaTroco.Login= v.un
		fmt.Println("Teste Usuario",k, v.un)
	}
	
	if req.Method == http.MethodPost {
		fltStr,err := strconv.ParseFloat(req.FormValue("valor"),64)
		if(err != nil){
			http.Error(w, "Não possível realizar ", http.StatusForbidden)
			return
		}
		EntradaTroco.Valor = fltStr
		modelos.CadEntradaTroco(EntradaTroco)		
	}

	p1 := TodoPageData{EntradaTroco : modelos.GetEntradaTrocoAll()}
	tpl.ExecuteTemplate(w, "entradatroco.gohtml", p1)
}
func mesa(w http.ResponseWriter, req *http.Request) {
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}	
	var Mesa modelos.Mesa
	if req.Method == http.MethodPost {
		for _, v := range dbSessions {
			Mesa.Login= v.un
		}
		Mesa.Numero, _ = strconv.Atoi(req.FormValue("numero"))
		if modelos.GetMesaNumero(Mesa.Numero)==true{
			fmt.Println("Número ja cadastrado")
			http.Error(w, "Mesa com mesmo numero ativa ", http.StatusForbidden)
			return
		}

		Mesa.ValorTotal = 0
		Mesa.Ativa = true
		
		modelos.CadMesa(Mesa)		
	}
	//showSessions() 
	p1 := TodoPageData{Mesa : modelos.GetMesaAll()}
	tpl.ExecuteTemplate(w, "mesa.gohtml", p1)
}
func permissao(w http.ResponseWriter, req *http.Request) {	
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}


	var Permissao modelos.Permissao
	if req.Method == http.MethodPost {
		Permissao.Nome = req.FormValue("nome")
		Permissao.Descricao = req.FormValue("descricao")
		modelos.CadPermissao(Permissao)		
	}
	p1 := TodoPageData{Permissao : modelos.GetPermissaoALL()}
		tpl.ExecuteTemplate(w, "permissao.gohtml", p1)
}
func produtomesa(w http.ResponseWriter, req *http.Request) {
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	
	var ProdutoMesa modelos.ProdutoMesa
	if req.Method == http.MethodPost {
		for k, v := range dbSessions {
			ProdutoMesa.Login= v.un
			fmt.Println("Teste Usuario",k, v.un)
		}
		
		if s, err := strconv.ParseInt(req.FormValue("quantidade"), 10, 64); err == nil {
			ProdutoMesa.Quantidade = int(s)
		}
		if s, err := strconv.ParseInt(req.FormValue("mesa"), 10, 64); err == nil {
			ProdutoMesa.Mesa.ID = int(s)
		}

		ProdutoMesa.Produto.ID,_ = strconv.Atoi(req.FormValue("produto"))
		teste,_ := strconv.Atoi(req.FormValue("produto"))
		listaProduto := modelos.GetProdutoIdAll(teste)
	
		for _,dado := range listaProduto{
			ProdutoMesa.ValorUnidade = dado.ValorUnidade
			ProdutoMesa.ValorTotal = ProdutoMesa.ValorUnidade* float64(ProdutoMesa.Quantidade)
		}

		fmt.Println("Teste produto mesa",ProdutoMesa)

		modelos.CadProdutoMesa(ProdutoMesa)		
	}

	fmt.Println("Chegou aqui lista final")
	//showSessions() 
	p1 := TodoPageData{ProdutoMesa : modelos.GetProdutoMesaAll(),Mesa : modelos.GetMesaAll(),Produto : modelos.GetProdutoAll()}
	tpl.ExecuteTemplate(w, "produtomesa.gohtml", p1)
}
func produtotipo(w http.ResponseWriter, req *http.Request) {
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	
	fmt.Println("Entrou no Cadastro de tipo de produtos")
	var ProdutoTipo modelos.ProdutoTipo
	if req.Method == http.MethodPost {
		ProdutoTipo.Nome = req.FormValue("nome")
		ProdutoTipo.Descricao = req.FormValue("descricao")
		modelos.CadProdutoTipo(ProdutoTipo)		
	}
	//showSessions() 
	p1 := TodoPageData{ProdutoTipo : modelos.GetProdutoTipoAll()}
	tpl.ExecuteTemplate(w, "produtotipo.gohtml", p1)
}
func produto(w http.ResponseWriter, req *http.Request) {
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
		
	var Produto modelos.Produto
	if req.Method == http.MethodPost {

		Produto.Nome = req.FormValue("nome")
		Produto.Descricao = req.FormValue("descricao")
		Produto.Quantidade,_ = strconv.Atoi(req.FormValue("quantidade"))

		if s, err := strconv.ParseInt(req.FormValue("produtotipo"), 10, 64); err == nil {
			Produto.ProdutoTipo.ID = int(s)
		}
		Produto.ValorUnidade,_= strconv.ParseFloat(req.FormValue("valorunidade"),64)
	
		modelos.CadProduto(Produto)		
	}

	//showSessions() 
	p1 := TodoPageData{Produto : modelos.GetProdutoAll(),ProdutoTipo : modelos.GetProdutoTipoAll()}
		tpl.ExecuteTemplate(w, "produto.gohtml", p1)
}
func sangria(w http.ResponseWriter, req *http.Request) {
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var Sangria modelos.Sangria
	if req.Method == http.MethodPost {

		for k, v := range dbSessions {
			Sangria.Login= v.un
			fmt.Println("Teste Usuario",k, v.un)
		}
	
		fltStr,err := strconv.ParseFloat(req.FormValue("valor"),64)
		if(err != nil){
			http.Error(w, "Não possível realizar ", http.StatusForbidden)
			return
		}
		Sangria.Valor = fltStr
		modelos.CadSangria(Sangria)		
	}
	//showSessions() 
	p1 := TodoPageData{Sangria : modelos.GetSangriaAll()}
	tpl.ExecuteTemplate(w, "sangria.gohtml", p1)
}
func delusuario(w http.ResponseWriter, req *http.Request){

	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if s, err := strconv.ParseInt(req.FormValue("id"), 10, 64); err == nil {
	modelos.DeletarUsuario(int(s))
	}else{
	}
	http.Redirect(w, req, "/usuario", http.StatusMovedPermanently)
}
func altusuario(w http.ResponseWriter, req *http.Request){

	fmt.Println("AltAusuairo")
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var Usua modelos.Usuario
	// converte string em int
	idUs, _ := strconv.Atoi(req.FormValue("id"))
	//busca os usuario por id
	ok,usuarios :=	modelos.GetUsuarioID(idUs)

	//verifica se a usuarios na 
	if ok == false {
		http.Error(w, "Operação não Realizada, erro ao alterar o usuario", http.StatusForbidden)
		return
	}
	//popula struct
	for _, dado := range usuarios {
		Usua = dado
	}
	
	Usua.Senha = nil

	p1 := TodoPageData{Usuarios : modelos.GetUsuarioAll(),Permissao : modelos.GetPermissaoALL(),Usuario :Usua}
	tpl.ExecuteTemplate(w, "altusuario.gohtml", p1)
	//tpl.ExecuteTemplate(w, "usuario.gohtml", Usuario)
}
func usuario(w http.ResponseWriter, req *http.Request) {
	
	// verifica se o usuario esta logado
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
		
	var Usuario modelos.Usuario
	if req.Method == http.MethodPost {
	
		id := req.FormValue("id")
		un := req.FormValue("nome")
		l := req.FormValue("login")
		p := req.FormValue("senha")

		//converte campo do formulario em int
		Usuario.ID ,_ = strconv.Atoi(id)

		//verifica se campor login esta vazio, senão busta o login
		if(l!=""){
			ok,arr := modelos.GetUsuarioLogin(l)
		
		// se o resultado da busca for true
		if(ok==true){
			for _,dados := range arr{
	
				if(dados.ID == Usuario.ID){
					ok = false
				}
			}
		}
		
		if(ok == true){
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return

			}else{
				
				if s, err := strconv.ParseInt(req.FormValue("permissao"), 10, 64); err == nil {
					Usuario.Permissao.ID = int(s)
				}

				//criando a session		
				sID, _ := uuid.NewV4()
				c := &http.Cookie{
					Name:  "session",
					Value: sID.String(),
				}
				c.MaxAge = sessionLength
				http.SetCookie(w, c)
				dbSessions[c.Value] = session{un, time.Now()}
				
				//gerar hash da senha
				//sum := sha256.Sum256([]byte())
				bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
				if err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
			}
			//popula a struct 
			Usuario.ID ,_ = strconv.Atoi(id)
			Usuario.Nome=un
			Usuario.Login = l
			Usuario.Senha = bs
			dbUsers[un] = Usuario
		
			modelos.CadUsuario(Usuario)
			}
		}
	}

	//Adicionada a struct TodoPageData os array 
	p1 := TodoPageData{Usuarios : modelos.GetUsuarioAll(),Permissao : modelos.GetPermissaoALL()}
	//Chama o template passando a pagina e a struct
	tpl.ExecuteTemplate(w, "usuario.gohtml", p1)
}
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
//Codígo para gerar HASH
//se quiser aumetar a segunra basta alterar o tipo do "sha256.New"
func getCode(data string,informacao string) string {
	h := hmac.New(sha256.New, []byte(informacao))
	//io.WriteString(h, data)
	fmt.Println("Testes", h)

	return fmt.Sprintf("%x", h.Sum(nil))
}
func signup(w http.ResponseWriter, req *http.Request) {
		var u modelos.Usuario
	
	//metado de submissão formulario
	if req.Method == http.MethodPost {
	// pegar dados do formulario

		un := req.FormValue("nome")
		l := req.FormValue("login")
		p := req.FormValue("senha")
		
		if s, err := strconv.ParseInt(req.FormValue("permissao"), 10, 64); err == nil {
		u.Permissao.ID = int(s)
		}

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		
		//gerar hash da senha
		//sum := sha256.Sum256([]byte())

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}


		fmt.Println("Senha :", p)
		fmt.Println("Senha encr",bs)



		u = modelos.Usuario{0,un,l, bs, u.Permissao}
		dbUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}
func login(w http.ResponseWriter, req *http.Request) {
		
	
	if req.Method == http.MethodPost {
		var u modelos.Usuario
		un := req.FormValue("login")
		p := req.FormValue("senha")
	
		//verficica se a algum usuario com o login inserido
		logi,arr := modelos.GetUsuarioLogin(un)
		if logi == false {
			http.Error(w, "Senha e Login Inválidos", http.StatusForbidden)
			return
		}
				
		fmt.Println("ok", un)
		fmt.Println("dbUsers", dbUsers[un])
		fmt.Println("u", u)
	

		for i, dado := range arr {
		//popula a struct Usuario
		u=dado
		fmt.Println("Informação", dado)
		fmt.Println("Indice", i)
		u, ok := dbUsers[un]

		fmt.Println("Dados =", u)
		fmt.Println("Dados OKc=", ok)
	}
		//comparação do hash cadastrado no banco, com o hash da senha que foi digitada
		err := bcrypt.CompareHashAndPassword(u.Senha, []byte(p))
		if err != nil {
			http.Error(w, "Login ou senha, inválidos", http.StatusForbidden)
			return
		}
		// criar a sessao
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
	
		u.Senha = nil
		http.Redirect(w, req, "/usuario", http.StatusMovedPermanently)
	
		//imprimir dados da session
		showSessions()
		return
	}
	showSessions()
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
 
}
func logout(w http.ResponseWriter, req *http.Request) {
		//verificar se o usuario esta logado
	if (alreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	
	//deleta a session
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		fmt.Println("Entrou aqui")
		go cleanSessions()
	}

	//faz o redirecinamento
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
func authorized(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// code before
		if !alreadyLoggedIn(w, r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, r)
		// code after
	})
}







 


