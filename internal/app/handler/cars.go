package handler

import (
	"context"

	"gofiber-app/internal/app/db"
	"gofiber-app/internal/app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCars(c *fiber.Ctx) error {
	var cars []models.Car

	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var car models.Car
		if err := cursor.Decode(&car); err != nil {
			return err
		}
		cars = append(cars, car)
	}

	return c.JSON(cars)
}

func CreateCar(c *fiber.Ctx) error {
	car := new(models.Car)

	if err := c.BodyParser(car); err != nil {
		return err
	}

	if car.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Car body cannot be empty"})
	}

	insertResult, err := db.Collection.InsertOne(context.Background(), car)
	if err != nil {
		return err
	}

	car.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(car)
}

func UpdateCar(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	// update car Name
	var updateData struct {
		Name string `json:"name"`
	}

	// parse the request body
	if err := c.BodyParser(&updateData); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	// check car_anme is provided in the request body
	if updateData.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Car name cannot be empty"})
	}

	// define and filter the update objects
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"name": updateData.Name}}

	// set the bolleand value true 
	// update := bson.M{"$set": bson.M{"car_price": 2345}}


	_, err = db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

func DeleteCar(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	_, err = db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
