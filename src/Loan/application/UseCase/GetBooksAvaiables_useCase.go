package UseCase

import "ArquitecturaHexagonal_02/src/Loan/domain"

type GetBooksAvailables struct {
	book domain.BookRepository
}

func NewGetBooksAvailableUseCase(book domain.BookRepository) *GetBooksAvailables {
	return &GetBooksAvailables{book: book}
}

func (useCase *GetBooksAvailables) Execute() ([]domain.Book, error) {
	return useCase.book.GetBooks()
}

