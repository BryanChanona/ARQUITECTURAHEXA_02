package UseCase

import "ArquitecturaHexagonal_02/src/Loan/domain"

type BooksAvailable struct {
	book domain.BookRepository
}

func NewSaveBookAvailableUseCase(book domain.BookRepository) *BooksAvailable {
	return &BooksAvailable{book: book}
}
func (useCase *BooksAvailable) Execute(book domain.Book) error{
	book.Available = true
	return useCase.book.Save(book)

}