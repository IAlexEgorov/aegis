package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"k8s/packages/aegis"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		var id string = r.FormValue("id")
		var name string = r.FormValue("name")
		var namespace string = r.FormValue("namespace")

		var aegis aegis.Aegis = aegis.Aegis{Id: id, Name: name, Namespace: namespace}
		go aegis.CreateProject()
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
