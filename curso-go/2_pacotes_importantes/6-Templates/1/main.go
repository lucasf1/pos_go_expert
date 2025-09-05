package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	curso := Curso{"Go", 40}
	tpl := template.New("CursoTemplate")
	tpl, _ = tpl.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := tpl.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
