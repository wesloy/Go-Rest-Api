package produtoRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meunegocio.com.br/api/db"
)

func Delete(id string) (deletedCount int64, err error) {

	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/deleteOne/

	client, ctx, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer context.WithCancel(ctx)

	collection := client.Database("estoque").Collection("produtos")

	idFilter, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": idFilter}
	result, err := collection.DeleteOne(ctx, filter)

	return result.DeletedCount, err

}
