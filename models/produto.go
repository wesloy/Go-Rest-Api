package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Produto struct {
	Id         primitive.ObjectID `bson:"_id, omitempty"`
	Nome       string             `json:"nome"  bson:"nome"`
	Marca      string             `json:"marca"  bson:"marca"`
	TipoVolume string             `json:"tipoVolume"  bson:"tipoVolume"`
	QtdeVolume float32            `json:"qtdeVolume"  bson:"qtdeVolume"`
	Validade   time.Time          `json:"validade"   bson:"validade"` // Golang suporta ISO 8601
}

// lista de produtos
type Produtos []*Produto

// omitempty omite o campo na serialização ou deserialização se estiver nulo
// bson é o mapeameto para o mongoDb
// _id é um campo obrigatório, já que o mongoDb sempre gera esse valor
