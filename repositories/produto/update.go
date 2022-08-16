package produtoRepository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meunegocio.com.br/api/db"
	"meunegocio.com.br/api/models"
)

func Update(id string, produto models.Produto) (modifiedCount int64, err error) {

	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/updateOne/

	client, ctx, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer context.WithCancel(ctx)
	collection := client.Database("estoque").Collection("produtos")

	idFilter, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": idFilter}
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "nome", Value: produto.Nome},
			{Key: "marca", Value: produto.Marca},
			{Key: "tipoVolume", Value: produto.TipoVolume},
			{Key: "qtdeVolume", Value: produto.QtdeVolume},
			{Key: "Validade", Value: produto.Validade},
		}}}

	result, err := collection.UpdateOne(ctx, filter, update)
	fmt.Println(result.ModifiedCount)

	return result.ModifiedCount, err

}

func UpdateAll(_key string, _value string, produto models.Produto) (modifiedCount int64, err error) {

	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/updateOne/

	client, ctx, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer context.WithCancel(ctx)
	collection := client.Database("estoque").Collection("produtos")

	filter := bson.M{_key: _value}

	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "nome", Value: produto.Nome},
			{Key: "marca", Value: produto.Marca},
			{Key: "tipoVolume", Value: produto.TipoVolume},
			{Key: "qtdeVolume", Value: produto.QtdeVolume},
			{Key: "Validade", Value: produto.Validade},
		}}}

	result, err := collection.UpdateMany(ctx, filter, update)
	fmt.Println(result.ModifiedCount)

	return result.ModifiedCount, err

}
