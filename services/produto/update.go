package produtoService

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"meunegocio.com.br/api/models"
	repository "meunegocio.com.br/api/repositories/produto"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
func Update(w http.ResponseWriter, r *http.Request) {

	// strconv é interno do GO e converte string para inteiro
	// chi é um pacote que auxilia a captura de informações das rotas http
	id := chi.URLParam(r, "id")

	// transferindo a informação vinda como Json para uma model
	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %e", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// atualizando no banco de dados...
	rows, err := repository.Update(chi.URLParam(r, "id"), produto)
	if err != nil {
		log.Printf("Erro ao fazer atualização do registro: %e", err) // print do erro
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// validando se teve apenas 1 registros atualizados
	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	// Mensagem final caso não tenha erros
	resp := map[string]any{
		"Status":  fmt.Sprintf("%d - %s", http.StatusOK, http.StatusText(http.StatusOK)),
		"Message": fmt.Sprintf("Informação atualizada com sucesso! ID: %s", id),
	}

	// incluíndo cabeçalho e a validação dentro do Json de resposta
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Exemplo URL, chamada:
// http://localhost:9000/marca-MINERVA
// http://localhost:9000/nome-Sabão liquido

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
func UpdateAll(w http.ResponseWriter, r *http.Request) {

	// chi é um pacote que auxilia a captura de informações das rotas http
	k := chi.URLParam(r, "k")
	v := chi.URLParam(r, "v")

	fmt.Printf("key: %s \n value:%s", k, v)

	// transferindo a informação vinda como Json para uma model
	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %e", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Observação:
	// Todos os documentos que derem match com a chave/valor serão atualziadas
	// se omitida par de chave/valor será setado vazio para o documento

	// atualizando no banco de dados...
	rows, err := repository.UpdateAll(k, v, produto)
	if err != nil {
		log.Printf("Erro ao fazer atualização do registro: %e", err) // print do erro
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Mensagem final caso não tenha erros
	resp := map[string]any{
		"Status":  fmt.Sprintf("%d - %s", http.StatusOK, http.StatusText(http.StatusOK)),
		"Message": fmt.Sprintf("Informação atualizada com sucesso! Total de registros: %d", rows),
	}

	// incluíndo cabeçalho e a validação dentro do Json de resposta
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
