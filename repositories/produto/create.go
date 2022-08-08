package produtoRepository

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"meunegocio.com.br/api/db"
	"meunegocio.com.br/api/models"
)

func Create(produto models.Produto) (insertedID primitive.ObjectID, err error) {

	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/insertOne/

	produto.Id = primitive.NewObjectID()

	client, ctx, err := db.OpenConnection()
	if err != nil {
		return
	}
	collection := client.Database("estoque").Collection("produtos")

	result, err := collection.InsertOne(ctx, produto)
	fmt.Println(result.InsertedID)

	insertedID, err = primitive.ObjectIDFromHex(fmt.Sprintf("%s", result.InsertedID))

	return insertedID, err

}
