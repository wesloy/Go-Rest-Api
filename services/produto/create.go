package produtoService

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Create(w http.ResponseWriter, r *http.Request) {

	// transferindo a informação vinda como Json para uma model
	var produto models.Produto

	// body, _ := ioutil.ReadAll(r.Body)
	// fmt.Printf(string(body))

	err := json.NewDecoder(r.Body).Decode(&produto)

	if err != nil {
		log.Printf("Erro ao fazer decode do Json: %v", err) // print do erro
		// informando o erro do servidor
		// sintaxe: onde vai escrever, mensagem, código da mensagem
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// inserindo no banco de dados...
	result, err := repository.Create(&produto)

	var resp map[string]any // variável cria uma tupla (Json)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		resp = map[string]any{
			"Status":  fmt.Sprintf("%s - %s", http.StatusText(http.StatusInternalServerError), http.StatusText(http.StatusInternalServerError)),
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {

		resp = map[string]any{
			"Status":  fmt.Sprintf("%d - %s", http.StatusCreated, http.StatusText(http.StatusCreated)),
			"Message": fmt.Sprintf("Informação inserida com sucesso! ID: %v", result.InsertedID.(primitive.ObjectID).Hex()),
		}
	}

	// incluíndo cabeçalho e a validação dentro do Json de resposta
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
