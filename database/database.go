package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Users  *mongo.Collection
	Movies *mongo.Collection
}

func Initialise() (*Database, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("database: No .env file found")
		// log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		return nil, fmt.Errorf("Database: You must set your 'MONGODB_URI' environmental variable")
		// log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	ctx := context.TODO()

	opt := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("Database: Cannot connect to DB")
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// Check the connection
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db := client.Database("newRatingMovies")
	usersCollection := db.Collection("users")
	moviesCollection := db.Collection("movies")

	fmt.Println("Connected to database")

	return &Database{
		Users:  usersCollection,
		Movies: moviesCollection,
	}, nil
}
