package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD */
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user-admin:twittorAdmin@twittole.naoor.mongodb.net/twittoleDB?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar la BD */
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		log.Println("error en la conexion ojo con el ping")
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
