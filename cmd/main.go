package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	PageTitle string
}

var tmpl = template.Must(template.ParseFiles("views/hello.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	data := Data{PageTitle: "Hello Live Reload"}
	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", Home)

	fmt.Printf("Server Listening at http://localhost:8001\n")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
}
