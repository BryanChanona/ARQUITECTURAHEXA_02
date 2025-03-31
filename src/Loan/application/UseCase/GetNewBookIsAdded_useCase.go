package UseCase

import "ArquitecturaHexagonal_02/src/Loan/domain"

type BookIsAddedUc struct {
	db domain.BookRepository
}

func NewNewBookIsAddedUc(db domain.BookRepository)*BookIsAddedUc{
	return &BookIsAddedUc{db: db}
}

func (useCase *BookIsAddedUc)Execute()(bool,[]domain.Book,error){
	return useCase.db.GetNewBookIdAdded()
}