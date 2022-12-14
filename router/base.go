package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"meunegocio.com.br/api/config"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "meunegocio.com.br/api/docs" // docs is generated by Swag CLI, you have to import it.
	servicesProduto "meunegocio.com.br/api/services/produto"
)

func base() (r *chi.Mux) {

	// criando as rotas de navegação
	r = chi.NewRouter()

	// injetando o Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", config.GetServerPort())), //The url pointing to API definition
	))

	return nil
}
func produto(r *chi.Mux) *chi.Mux {

	r.Post("/", servicesProduto.Create)
	r.Put("/{id}", servicesProduto.Update)
	r.Put("/{k}-{v}", servicesProduto.UpdateAll)
	r.Delete("/{id}", servicesProduto.Delete)
	r.Get("/", servicesProduto.GetAll)
	r.Get("/{id}", servicesProduto.Get)

	return r

}

func CarregarRotas() {

	r := chi.NewRouter()

	// injetando o Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", config.GetServerPort())), //The url pointing to API definition
	))

	// Carregando as rotas de cada uma das funcionalidades da API
	// PRODUTO
	r.Post("/", servicesProduto.Create)
	r.Put("/{id}", servicesProduto.Update)
	r.Put("/{k}-{v}", servicesProduto.UpdateAll)
	r.Delete("/{id}", servicesProduto.Delete)
	r.Get("/", servicesProduto.GetAll)
	r.Get("/{id}", servicesProduto.Get)

	// montando o servidor para escutar as chamadas
	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
