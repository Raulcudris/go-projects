package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Raulcudris/go-react-crud/models"
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

	//Uso de los Cors(Permisos en los navegadores )
	app.Use(cors.New())

	//Direccion para la llamada a el Front
	app.Static("/", "./client/dist")

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		c.BodyParser(&user)
		//si la base de Datos con mongo conecta hacemos una consulta
		coll := client.Database("gomongodb").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	//Creacion de un Endpoint para la llamada a el Front
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		coll := client.Database("gomongodb").Collection("users")
		results, err := coll.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"data": users,
		})
	})

	//Listening Port
	app.Listen(":" + port)
	fmt.Println("Server on port " + port)
}
