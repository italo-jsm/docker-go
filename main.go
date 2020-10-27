package main

import (
	"net/http"
)

type appHandler struct{
	Message string
}

func (a *appHandler) ServeHTTP(w http.ResponseWriter, r * http.Request){
	w.Write([]byte(a.Message))
}

func main(){
	http.Handle("/app", &appHandler{Message: "Server online!"})
	http.ListenAndServe(":3000", nil)
}