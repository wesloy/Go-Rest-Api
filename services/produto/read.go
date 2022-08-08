package produtoService

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	repository "meunegocio.com.br/api/repositories/produto"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
func Get(w http.ResponseWriter, r *http.Request) {
	// strconv é interno do GO e converte string para inteiro
	// chi é um pacote que auxilia a captura de informações das rotas http

	// buscando no banco de dados...
	produto, err := repository.Get(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao buscar o registro  %e", err) // print do erro
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// incluíndo cabeçalho e a validação dentro do Json de resposta
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produto)
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
func GetAll(w http.ResponseWriter, r *http.Request) {

	produtos, err := repository.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	// incluíndo cabeçalho e o array de resposta, vazio ou preenchido
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)

}
