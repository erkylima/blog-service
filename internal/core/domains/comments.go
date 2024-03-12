package domains

type Comment struct {
	Id        string `json:"id" bson:"_id"`
	Body      string `json:"body" bson:"body"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeletedAt string `json:"deleted_at" bson:"deleted_at"`
	Author    string `json:"author" bson:"author"`
	PostId    string `json:"post_id" bson:"post_id"`
}
