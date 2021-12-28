package models

type Livro struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Titulo string `json:"titulo"`
}

var Livros []Livro
