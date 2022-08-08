package router

import (
	"github.com/go-chi/chi"
	servicesProduto "meunegocio.com.br/api/services/produto"
)

func Produto(r *chi.Mux) *chi.Mux {

	r.Post("/", servicesProduto.Create)
	r.Put("/{id}", servicesProduto.Update)
	r.Put("/{k}-{v}", servicesProduto.UpdateAll)
	r.Delete("/{id}", servicesProduto.Delete)
	r.Get("/", servicesProduto.GetAll)
	r.Get("/{id}", servicesProduto.Get)

	return r

}
