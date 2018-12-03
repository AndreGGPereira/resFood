package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/andreggpereira/resFood/modelos"
	"github.com/satori/go.uuid"
	
)

//GetUser pegar usuario
func getUser(w http.ResponseWriter, req *http.Request) modelos.Usuario {
	
	//tenta pegar o Cookie
	c, err := req.Cookie("session")
	// se não achou, cria novo
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// se usuario existe, retorna o usuario
	var u modelos.Usuario
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}
//Metado  serve para testar se o usuario esta logado
func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	fmt.Println(" C ", c)
	s, ok := dbSessions[c.Value]
	fmt.Println(" S ", s)
	fmt.Println(" OK ", ok)
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	fmt.Println(" OK Final",ok)

	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
//CleanSessions limpar Session
func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}
//ShowSessions imprimir dados da session
func showSessions() {
	fmt.Println("********")
	fmt.Println("Dados da Session")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}

