package ports

type Repository interface {
	Create(entity interface{}) (*[]byte, error)
	Read(slug string, entity interface{}) (*[]byte, error)
	Update(entity interface{}) (*[]byte, error)
	Delete(slug string, entity interface{}) error
	List(entity interface{}) ([]byte, error)
	ListByTag(tag string, entity interface{}) ([]byte, error)
	ListByCategory(category string, entity interface{}) ([]byte, error)
	ListByAuthor(author string, entity interface{}) ([]byte, error)
	ListByDate(date string, entity interface{}) ([]byte, error)
}
