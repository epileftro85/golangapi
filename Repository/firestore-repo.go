package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	entity "github.com/epileftro85/goapi/Entity"
	"google.golang.org/api/iterator"
)

const (
	PROJECT_ID string = "golang-b95ed"
	COLLECTION string = "posts"
)

type firestoreRepo struct{}

// Constructor
func NewFirestoreRepository() PostRepository {
	return &firestoreRepo{}
}

func (*firestoreRepo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Error on creating new firestore client %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(COLLECTION).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Error on adding new post %v", err)
		return nil, err
	}

	return post, nil
}

func (*firestoreRepo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Error on creating new firestore client %v", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	iter := client.Collection(COLLECTION).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error on iterate all posts %v", err)
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
