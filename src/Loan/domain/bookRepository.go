package domain

type BookRepository interface {
	Save(book Book) error
	GetBooks()([]Book, error)
	GetNewBookIdAdded()(bool,[]Book, error)
}
