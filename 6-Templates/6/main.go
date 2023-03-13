package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		Curso{Nome: "C#", CargaHoraria: 10},
		Curso{Nome: "Go", CargaHoraria: 20},
	})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8282", nil)
}
