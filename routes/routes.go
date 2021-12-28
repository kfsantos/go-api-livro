package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kfsantos/livros/controllers"
	"github.com/kfsantos/livros/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/livros", controllers.ListarLivros).Methods("Get")
	r.HandleFunc("/livros/{id}", controllers.BuscarLivrosPorId).Methods("Get")
	r.HandleFunc("/livros/novo", controllers.CriarLivro).Methods("Post")
	r.HandleFunc("/livros/{id}", controllers.EditaLivro).Methods("Put")
	r.HandleFunc("/livros/{id}", controllers.RemoverLivro).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
