package domains

type Category struct {
	Id   string `json:"id" bson:"_id"`
	Slug string `json:"slug" bson:"slug"`
	Name string `json:"name" bson:"name"`
}
