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

	app, err := firebase.NewApp(ctx, nil, getClientOption())

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func getClientOption() option.ClientOption {

	if option := getClientOptionWithJsonString(); option != nil {
		return option
	}

	if option := getClientOptionWithFile(); option != nil {
		return option
	}

	return nil
}

func getClientOptionWithFile() option.ClientOption {

	filename := "serviceAccount.json"

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return option.WithCredentialsFile(filename)
}

func getClientOptionWithJsonString() option.ClientOption {

	if env := os.Getenv("GOOGLE_SERVICE_ACCOUNT"); env != "" {
		return option.WithCredentialsJSON([]byte(env))
	}

	return nil
}
