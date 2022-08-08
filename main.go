// https://www.youtube.com/watch?v=MD7b-iQMC24
// https://www.youtube.com/watch?v=_4JpGnoh0Pg

package main

import (
	"meunegocio.com.br/api/config"
	router "meunegocio.com.br/api/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {

	// buscando configurações e se não achar parando a execução
	err := config.Load()
	if err != nil {
		panic(err)
	}

	router.CarregarRotas()

	// // criando as rotas de navegação
	// r := chi.NewRouter()

	// r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost:9000/swagger/doc.json"), //The url pointing to API definition
	// ))

	// r.Post("/", servicesProduto.Create)
	// r.Put("/{id}", servicesProduto.Update)
	// r.Put("/{k}-{v}", servicesProduto.UpdateAll)
	// r.Delete("/{id}", servicesProduto.Delete)
	// r.Get("/", servicesProduto.GetAll)
	// r.Get("/{id}", servicesProduto.Get)

	// // montando o servidor para escutar as chamadas
	// http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
