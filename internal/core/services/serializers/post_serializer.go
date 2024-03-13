package serializers

import (
	"encoding/json"

	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/pkg/errors"
)

type postSerializer struct{}

func NewPostSerializer() *postSerializer {
	return &postSerializer{}
}

func (co *postSerializer) Encode(input domains.Post) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.post.encode")
	}
	return rawMsg, nil

}

func (co *postSerializer) Decode(input []byte) (*domains.Post, error) {
	post := &domains.Post{}
	if err := json.Unmarshal(input, &post); err != nil {
		return nil, errors.Wrap(err, "serializer.post.decode")
	}
	return post, nil
}
