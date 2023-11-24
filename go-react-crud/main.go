package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//Habilitacion de Puerto cuando vaya a Produccion
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//Llamada de Fiber (Framework para la creacion del servidor)
	app := fiber.New()

	//Llamada a la base de Datos con mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb"))
	//si la base de Datos con mongo no conecta
	if err != nil {
		panic(err)
	}
	//si la base de Datos con mongo conecta hacemos una consulta
	coll := client.Database("gomongodb").Collection("users")
	coll.InsertOne(context.TODO(), bson.D{{
		Key:   "name",
		Value: "Raul Cudris",
	}})

	//Uso de los Cors(Permisos en los navegadores )
	app.Use(cors.New())

	//Direccion para la llamada a el Front
	app.Static("/", "./client/dist")

	//Creacion de un Endpoint para la llamada a el Front
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Usuarios desde el backend"})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Creando Usuario",
		})
	})

	//Listening Port
	app.Listen(":" + port)
	fmt.Println("Server on port " + port)
}
