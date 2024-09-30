package main

import (
	"ascii-art/functions"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"strings"
)

type Data struct {
	Str string
	Banner string
	Res string
	A	template.HTML
}
type Test struct {
	Name string
	Age int
}
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Mostafa")
// }

func processHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var  d Data
	d.Str = r.FormValue("data")
	d.Banner = r.FormValue("banner")
	log.Println("Received Data:", d.Str)
	d.Str = strings.ReplaceAll(d.Str, "\r\n", "\\n")
	d.Str = strings.ReplaceAll(d.Str, "\n", "\\n")
	d.Res = function.TraitmentData(d.Banner, d.Str)
	d.A = template.HTML(strings.ReplaceAll(d.Res, "\n", "<br>"))
	t.Execute(w, d)
	// if err := t.Execute(w, d); err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    // }
	// io.WriteString(w, res)
}



func ageHandler(w http.ResponseWriter, r *http.Request) {
	// a := Test{Name: "Mostafa", Age: 27}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error()/*convert the error into string*/, http.StatusInternalServerError)
	}
	t.Execute(w, "")
}


func main() {
	// http.HandleFunc("/", indexHandler)
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	http.HandleFunc("/home", ageHandler)
	http.HandleFunc("/process", processHandler)
	http.ListenAndServe(":8080", nil)
}