package ports

type Repository interface {
	Create(entity interface{}) (string, error)
	Read(slug string, entity interface{}) (interface{}, error)
	Update(slug string, entity interface{}) (interface{}, error)
	Delete(slug string, entity interface{}) error
	List(entity interface{}) (interface{}, error)
	ListByTag(tag string, entity interface{}) (interface{}, error)
	ListByCategory(category string, entity interface{}) (interface{}, error)
	ListByAuthor(author string, entity interface{}) (interface{}, error)
	ListByDate(date string, entity interface{}) (interface{}, error)
}
