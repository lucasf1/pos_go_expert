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

	// tpl := template.New("CursoTemplate")
	// tpl, err := tpl.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	// if err != nil {
	// 	panic(err)
	// }

	// template.must é uma função que cria um novo template e o analisa em uma única etapa.
	// Se ocorrer um erro durante a análise, o programa será encerrado com um pânico
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))

	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}
