package entities

type Tag struct {
	Slug string `json:"slug" bson:"slug"`
	Name string `json:"name" bson:"name"`
}
