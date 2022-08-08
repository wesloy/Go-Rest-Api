package db

import (
	"context"
	"fmt"

	"meunegocio.com.br/api/config"

	// coloca o _ para que o GO não remova uma dependencia se não tiver sendo utilizada
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (client *mongo.Client, ctx context.Context, err error) {

	conf := config.GetDB()
	sc := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		conf.User, conf.Pass, conf.Host, conf.Port)

	ctx = context.Background()

	// conexão Mongo
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(sc))

	// // testando conexão
	// err = client.Ping(ctx, options.Client().ReadPreference)
	// if err == nil {
	// 	println("Conexão realizada com sucesso!")
	// }

	return

}
