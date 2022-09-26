package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer .html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()
	//var a aegis.Aegis = aegis.Aegis{Id: "test1", Name: "test", Namespace: "test"}
	//db := a.ConnectToDB("root", "root", "Aegis", 8889)
	//a.CreateProject(db)
}
