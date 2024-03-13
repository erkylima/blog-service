package serializers

import (
	"encoding/json"

	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/pkg/errors"
)

type categorySerializer struct{}

func NewCategorySerializer() *categorySerializer {
	return &categorySerializer{}
}

func (ca *categorySerializer) Encode(input domains.Category) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.category.encode")
	}
	return rawMsg, nil

}

func (ca *categorySerializer) Decode(input []byte) (*domains.Category, error) {
	dockercompose := &domains.Category{}
	if err := json.Unmarshal(input, &dockercompose); err != nil {
		return nil, errors.Wrap(err, "serializer.category.decode")
	}
	return dockercompose, nil
}
