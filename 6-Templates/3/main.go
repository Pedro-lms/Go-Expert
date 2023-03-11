package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"C#", 10},
	})
	if err != nil {
		panic(err)
	}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
}
