package domains

import (
	comment_domain "github.com/erkylima/posts-service/internal/comments/domains"
	"github.com/erkylima/posts-service/internal/shared/entities"
)

type Post struct {
	Id           string                   `json:"id" bson:"_id"`
	Slug         string                   `json:"slug" bson:"slug"`
	Title        string                   `json:"title" bson:"title"`
	Description  string                   `json:"description" bson:"description"`
	Body         string                   `json:"body" bson:"body"`
	CreatedAt    string                   `json:"created_at" bson:"created_at"`
	UpdatedAt    string                   `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt    string                   `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	PublishedAt  string                   `json:"published_at,omitempty" bson:"published_at,omitempty"`
	Author       string                   `json:"author,omitempty" bson:"author,omitempty"`
	Tags         []entities.Tag           `json:"tags,omitempty" bson:"tags,omitempty"`
	CommentCount int                      `json:"comment_count,omitempty" bson:"comment_count,omitempty"`
	Comments     []comment_domain.Comment `json:"comments,omitempty" bson:"comments,omitempty"`
}
