package repositories

import "ArquitecturaHexagonal_02/src/Loan/domain"

type IEventPublisher interface {
	PublishEvent(book domain.Book) error
}