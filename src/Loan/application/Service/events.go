package service

import (
	repositories "ArquitecturaHexagonal_02/src/Loan/application/repositories"
	dominio "ArquitecturaHexagonal_02/src/Loan/domain"
)

type Event struct {
	rabbit repositories.IEventPublisher
}

func NewEvent(rabbit repositories.IEventPublisher)*Event{
	return &Event{rabbit:rabbit}
}

func (event *Event) Execute(book dominio.Book)(err error){
	return event.rabbit.PublishEvent(book)
}