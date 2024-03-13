package ports

import "github.com/erkylima/posts-service/internal/core/domains"

type CommentService interface {
	CreateComment(comment domains.Comment) (*domains.Comment, error)
	ReadComment(id string) (*domains.Comment, error)
	UpdateComment(comment domains.Comment) (*domains.Comment, error)
	DeleteComment(id string) error
	ListComments() ([]domains.Comment, error)
	ListCommentsByPost(postId string) ([]domains.Comment, error)
	ListCommentsByAuthor(author string) ([]domains.Comment, error)
	ListCommentsByTag(tag string) ([]domains.Comment, error)
}

type CommentRepository interface {
	CreateComment(comment domains.Comment) (*domains.Comment, error)
	ReadComment(id string) (*domains.Comment, error)
	UpdateComment(comment domains.Comment) (*domains.Comment, error)
	DeleteComment(id string) error
	ListComments() ([]domains.Comment, error)
	ListCommentsByPost(postId string) ([]domains.Comment, error)
	ListCommentsByAuthor(author string) ([]domains.Comment, error)
	ListCommentsByTag(tag string) ([]domains.Comment, error)
}
