package produtoService

import (
	"encoding/json"
	"fmt"
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
func Delete(w http.ResponseWriter, r *http.Request) {

	// strconv é interno do GO e converte string para inteiro
	// chi é um pacote que auxilia a captura de informações das rotas http
	id := chi.URLParam(r, "id")
	// deletando no banco de dados...
	rows, err := repository.Delete(id)
	if err != nil {
		log.Printf("Erro ao fazer deleção do registro: %e", err) // print do erro
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// validando se teve apenas 1 registros atualizados
	if rows > 1 {
		log.Printf("Error: foram deletados %d registros", rows)
	}

	// Mensagem final caso não tenha erros
	resp := map[string]any{
		"Status":  fmt.Sprintf("%d - %s", http.StatusOK, http.StatusText(http.StatusOK)),
		"Message": fmt.Sprintf("Deletado registro com sucesso! ID: %s", id),
	}

	// incluíndo cabeçalho e a validação dentro do Json de resposta
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
