package domains

type Navigation struct {
	Slug string `json:"slug" bson:"slug"`
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
	Type string `json:"type" bson:"type"`
	Id   string `json:"id" bson:"_id"`
}
