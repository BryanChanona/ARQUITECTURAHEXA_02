package domain

type BookRepository interface {
	Save(book Book) error
}
