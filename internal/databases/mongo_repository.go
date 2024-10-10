package databases

import (
	"context"

	"github.com/Israel-Ferreira/shopping-api/pkg/interfaces"
	"github.com/Israel-Ferreira/shopping-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	DbConn *mongo.Client
}

func (mr MongoRepository) ListShoppings() ([]models.Shopping, error) {
	shoppingCollection := GetCollection(mr.DbConn, "shoppings", "shoppings")
	cur, err := shoppingCollection.Find(context.Background(), bson.D{})

	ctx := context.Background()

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var shoppings []models.Shopping

	for cur.Next(ctx) {
		var shopping models.Shopping

		if err := cur.Decode(&shopping); err != nil {
			return nil, err
		}

		shoppings = append(shoppings, shopping)
	}

	return shoppings, nil
}

func (mr MongoRepository) GetShoppingById(id string) (models.Shopping, error) {
	shoppingCollection := GetCollection(mr.DbConn, "shoppings", "shoppings")

	var shopping models.Shopping

	ctx := context.Background()

	if err := shoppingCollection.FindOne(ctx, bson.M{"id": id}).Decode(&shopping); err != nil {
		return shopping, err
	}

	return shopping, nil
}

func (mr MongoRepository) DeleteShoppingById(id string) error {
	shoppingCollection := GetCollection(mr.DbConn, "shoppings", "shoppings")

	ctx := context.Background()

	if _, err := shoppingCollection.DeleteOne(ctx, bson.M{"id": id}); err != nil {
		return err
	}

	return nil
}

func (mr MongoRepository) UpdateShopping(id string, dtoReq models.ShoppingRequest) error {
	return nil
}

func (mr MongoRepository) CreateShopping(dtoReq models.ShoppingRequest) error {
	shoppingCollection := GetCollection(mr.DbConn, "shoppings", "shoppings")

	newShopping := models.NewShopping(dtoReq)

	ctx := context.Background()

	if _, err := shoppingCollection.InsertOne(ctx, newShopping); err != nil {
		return err
	}

	return nil
}

func NewRepository(client *mongo.Client) interfaces.DbRepository {
	return MongoRepository{
		DbConn: client,
	}
}
