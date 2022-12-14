package connector

import (
	"comm-api/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func Init() {

	var connectionUri string = models.DB_URL
	var dbName = models.DB_NAME
	var collectionName = models.COLLECTION_NAME

	clientOption := options.Client().ApplyURI(connectionUri)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	//collection instance

	fmt.Println("Collection instance is ready")
}

func InsertOne(msg interface{}) error {

	inserted, err := collection.InsertOne(context.Background(), msg)

	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted one expense with ID:", inserted.InsertedID)
	return nil

}

func GetOne(id string) (models.Message, error) {

	msgId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": msgId}

	res := collection.FindOne(context.Background(), filter)
	var msg models.Message

	if res.Err() != nil {
		return msg, res.Err()
	}

	res.Decode(&msg)

	return msg, nil

}

func GetAll() ([]models.Message, error) {

	curr, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	defer curr.Close(context.Background())

	var messages []models.Message

	for curr.Next(context.Background()) {
		var message models.Message
		if err := curr.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil

}

func DeleteOne(id string) error {

	msgId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": msgId}

	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	fmt.Println("Successfully deleted number of messages: ", res.DeletedCount)
	return nil

}

func DeleteAll() error {

	curr, err := collection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return nil
	}

	fmt.Println("Successfully deleted number of messages: ", curr.DeletedCount)

	return nil
}

func UpdateOne(id string, msgbody string) error {

	msgId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": msgId}

	new_filter := bson.M{"$set": bson.M{"message": msgbody}}

	res, err := collection.UpdateOne(context.Background(), filter, new_filter)

	if err != nil {
		return err
	}

	// var msg models.Message

	fmt.Println("Seccessfully updates number of msg: ", res.ModifiedCount)

	return nil

}
