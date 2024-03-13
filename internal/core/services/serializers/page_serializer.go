package serializers

import (
	"encoding/json"

	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/pkg/errors"
)

type pageSerializer struct{}

func NewPageSerializer() *pageSerializer {
	return &pageSerializer{}
}

func (co *pageSerializer) Encode(input domains.Page) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.page.encode")
	}
	return rawMsg, nil

}

func (co *pageSerializer) Decode(input []byte) (*domains.Page, error) {
	page := &domains.Page{}
	if err := json.Unmarshal(input, &page); err != nil {
		return nil, errors.Wrap(err, "serializer.page.decode")
	}
	return page, nil
}
