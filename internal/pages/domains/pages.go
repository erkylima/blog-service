package domains

import "github.com/erkylima/posts-service/internal/shared/entities"

type Page struct {
	Id        string         `json:"id,omitempty" bson:"_id,omitempty"`
	Slug      string         `json:"slug" bson:"slug"`
	Title     string         `json:"title" bson:"title"`
	Url       string         `json:"url" bson:"url"`
	Type      string         `json:"type" bson:"type"`
	Tags      []entities.Tag `json:"tags" bson:"tags"`
	Body      string         `json:"body" bson:"body"`
	CreatedAt string         `json:"created_at" bson:"created_at"`
	UpdatedAt string         `json:"updated_at" bson:"updated_at"`
	DeletedAt string         `json:"deleted_at" bson:"deleted_at"`
}
