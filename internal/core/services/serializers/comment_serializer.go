package serializers

import (
	"encoding/json"

	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/pkg/errors"
)

type commentSerializer struct{}

func NewCommentSerializer() *commentSerializer {
	return &commentSerializer{}
}

func (co *commentSerializer) Encode(input domains.Comment) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.comment.encode")
	}
	return rawMsg, nil

}

func (co *commentSerializer) Decode(input []byte) (*domains.Comment, error) {
	comment := &domains.Comment{}
	if err := json.Unmarshal(input, &comment); err != nil {
		return nil, errors.Wrap(err, "serializer.comment.decode")
	}
	return comment, nil
}
