package main

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Transaction struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string
	Amount     float32
	Categories []string
}

func NewDatabase(ctx context.Context, url string) (database, error) {
	opts := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return database{}, err
	}

	db := database{
		ctx:        ctx,
		client:     client,
		Collection: client.Database("mydb").Collection("transactions"),
	}

	return db, nil
}

type database struct {
	ctx        context.Context
	client     *mongo.Client
	Collection *mongo.Collection
}

func (d *database) InsertTransaction(t *Transaction) error {
	t.ID = primitive.NewObjectID()

	response, err := d.Collection.InsertOne(d.ctx, t)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func (d *database) GetTransaction(id string) (Transaction, error) {
	parsedId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{
		"_id": parsedId,
	}

	transactions, err := d.GetTransactions(query)
	if err != nil {
		return Transaction{}, err
	}

	if len(transactions) > 1 {
		return Transaction{}, errors.New("multiple results got")
	}

	if len(transactions) < 1 {
		return Transaction{}, errors.New("no results found")
	}

	return transactions[0], nil
}

func (d *database) GetTransactions(query bson.M) ([]Transaction, error) {
	cursor, err := d.Collection.Find(d.ctx, query)
	if err != nil {
		return []Transaction{}, err
	}

	var results []Transaction
	if err = cursor.All(d.ctx, &results); err != nil {
		return []Transaction{}, err
	}

	return results, nil
}

func (d *database) DeleteTransaction(id string) error {
	parsedId, _ := primitive.ObjectIDFromHex(id)
	r, err := d.Collection.DeleteOne(d.ctx, bson.M{"_id": parsedId})

	fmt.Println(r)

	if err != nil {
		return err
	}

	return nil
}

func (d *database) UpdateTransaction(id string, t *Transaction) error {
	var err error
	t.ID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	body := bson.D{{"$set", t}}
	r, err := d.Collection.UpdateOne(d.ctx, bson.M{"_id": t.ID}, body)
	if err != nil {
		return err
	}

	fmt.Println(r)

	return nil
}
