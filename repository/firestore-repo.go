package repository

import (
	"../entity"
	"cloud.google.com/go/firestore"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	PROJECT_ID      string = "reviews-api-db"
	COLLECTION_NAME string = "posts"
)

type repo struct {
}

//NewFirestoreRepository ...
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, PROJECT_ID, option.WithCredentialsFile("./reviews-api-db-firebase-adminsdk-610bc-253e8e69f6.json"))
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(COLLECTION_NAME).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed to add a new post to db: %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, PROJECT_ID, option.WithCredentialsFile("./reviews-api-db-firebase-adminsdk-610bc-253e8e69f6.json"))
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iter := client.Collection(COLLECTION_NAME).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			log.Fatalf("Failed to iterate lists of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
