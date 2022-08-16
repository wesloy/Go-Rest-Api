package produtoRepository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meunegocio.com.br/api/db"
	"meunegocio.com.br/api/models"
)

func Get(id string) (produto models.Produto, err error) {

	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/findOne/

	client, ctx, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer context.WithCancel(ctx)

	collection := client.Database("estoque").Collection("produtos")

	idFilter, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return produto, err
	}

	filter := bson.M{"_id": idFilter}

	err = collection.FindOne(ctx, filter).Decode(&produto)
	fmt.Println(produto)

	return

}

func GetAll() (produtos []models.Produto, err error) {

	// https://www.mongodb.com/docs/drivers/go/current/usage-examples/find/

	client, ctx, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer context.WithCancel(ctx)
	collection := client.Database("estoque").Collection("produtos")

	cur, err := collection.Find(ctx, bson.M{}) // bson.M é pra trazer todos
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &produtos)
	if err != nil {
		log.Fatal(err)
	}

	return

}
