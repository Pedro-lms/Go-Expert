package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Name}} - CargaHoraria: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso)
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horária {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
}
