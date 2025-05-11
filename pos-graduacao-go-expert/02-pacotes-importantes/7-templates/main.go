package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	curso := Atividade{
		Nome:               "Curso de Go",
		TempoParaFinalizar: 60}
	livro := Atividade{
		Nome:               "O Jogo Do Anjo",
		TempoParaFinalizar: 20,
	}

	tmp := template.New("AtividadeTemplate")
	tmp, _ = tmp.Parse("A atividade {{.Nome}} tem tempo de finalização de {{.TempoParaFinalizar}} horas.\n")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("AtividadeTemplate").
		Parse("A atividade {{.Nome}} tem tempo de finalização de {{.TempoParaFinalizar}} horas.\n"))
	err = t.Execute(os.Stdout, livro)
	if err != nil {
		panic(err)
	}

	// Utilizando o template com arquivo
	// tdoc := template.Must(template.New("template.html").ParseFiles("template.html"))
	// ou tdoc := template.Must(template.ParseFiles("template.html"))
	atividades := Atividades{
		curso,
		livro,
		{Nome: "Curso de Go Expert", TempoParaFinalizar: 240},
		{Nome: "A sombra do vento", TempoParaFinalizar: 30},
	}
	// Substitui os valores do template.html pelos valores de atividades
	// err = tdoc.Execute(os.Stdout, atividades)
	// if err != nil {
	// 	panic(err)
	// }

	// Servidor HTTP com template
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	err = tdoc.Execute(w, atividades)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// })
	// http.ListenAndServe(":8181", nil)

	templates := []string{
		"header.html",
		"template.html",
		"footer.html",
	}

	// tep := template.Must(template.New("template.html").ParseFiles(templates...))
	
	tep := template.New("template.html")
	// Mapeamento de funções
	tep.Funcs(template.FuncMap{
		"MM": Maiuscula,
	})
	tep = template.Must(tep.ParseFiles(templates...))
	err = tep.Execute(os.Stdout, atividades)
	if err != nil {
		panic(err)
	}
}

type Atividade struct {
	Nome               string
	TempoParaFinalizar int
}

type Atividades []Atividade

func Maiuscula(s string) string {
	return strings.ToUpper(s)
}
