package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
)

var ctx = context.Background()

type Firebase struct {
}

// FindByKey requires a collection and a key value to be found.
func (f Firebase) FindByKey(collection, key string) map[string]interface{} {

	client := getClient()
	defer client.Close()

	document, err := client.Collection(collection).Doc(key).Get(ctx)

	if err != nil {
		return nil
	}

	return document.Data()
}

func (f Firebase) Store(collection, key string, data map[string]interface{}) {
	client := getClient()
	defer client.Close()

	_, err := client.Collection(collection).Doc(key).Create(ctx, data)

	if err != nil {
		log.Fatalf("Failed adding %s to firebase: %v", key, err)
	}
}

func getClient() *firestore.Client {

	filename := "serviceAccount.json"

	var app *firebase.App
	var err error

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		app, err = firebase.NewApp(ctx, nil)
	} else {
		// Use a service account
		sa := option.WithCredentialsFile("serviceAccount.json")
		app, err = firebase.NewApp(ctx, nil, sa)
	}

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	return client
}
