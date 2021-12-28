package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kfsantos/livros/database"
	"github.com/kfsantos/livros/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Página inicial!!!")
}

func ListarLivros(w http.ResponseWriter, r *http.Request) {
	var livros []models.Livro
	database.DB.Find(&livros)
	json.NewEncoder(w).Encode(livros)
}

func BuscarLivrosPorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var livro models.Livro
	database.DB.First(&livro, id)
	json.NewEncoder(w).Encode(livro)
}

func CriarLivro(w http.ResponseWriter, r *http.Request) {
	var novoLivro models.Livro
	json.NewDecoder(r.Body).Decode(&novoLivro)
	database.DB.Create(&novoLivro)
	json.NewEncoder(w).Encode(novoLivro)
}

func EditaLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var livro models.Livro
	database.DB.First(&livro, id)
	json.NewDecoder(r.Body).Decode(&livro)
	database.DB.Save(&livro)
	json.NewEncoder(w).Encode(livro)
}

func RemoverLivro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var livro models.Livro
	database.DB.Delete(&livro, id)
	json.NewEncoder(w).Encode("livro com id: " + id + " foi  removido ")
}
