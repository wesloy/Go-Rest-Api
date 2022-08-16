package produtoRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"meunegocio.com.br/api/db"
	"meunegocio.com.br/api/models"
)

func Create(produto *models.Produto) (result *mongo.InsertOneResult, err error) {

	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/insertOne/

	produto.Id = primitive.NewObjectID()

	client, ctx, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer context.WithCancel(ctx)

	collection := client.Database("estoque").Collection("produtos")
	result, err = collection.InsertOne(ctx, produto)

	return

}
