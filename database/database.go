package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	// Client *mongo.Client
	Users  *mongo.Collection
	Movies *mongo.Collection
}

func Initialise() (*Database, error) {

	if err := godotenv.Load(); err != nil {
		return nil, errors.New("database: No .env file found")
		// log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		return nil, errors.New("Database: You must set your 'MONGODB_URI' environmental variable")
		// log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	ctx := context.TODO()

	opt := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, errors.New("Database: Cannot connect to DB")
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.New("Database: Lost the connection to DB")
	}

	db := client.Database("newRatingMovies")
	usersCollection := db.Collection("users")
	moviesCollection := db.Collection("movies")

	// defer func() {
	// 	if err := client.Disconnect(ctx); err != nil {
	// 		fmt.Println("Error:", err)
	// 	}
	// }()

	fmt.Println("Connected to database")

	return &Database{
		// Client: client,
		Users:  usersCollection,
		Movies: moviesCollection,
	}, nil
}

// func CloseConnection(client *mongo.Client) error {
// 	if client == nil {
// 		return errors.New("no connection to disconnect")
// 	}

// 	err := client.Disconnect(context.TODO())
// 	if err != nil {
// 		return err
// 	}

// 	// TODO optional you can log your closed MongoDB client
// 	fmt.Println("Connection to MongoDB closed.")

// 	return nil
// }
