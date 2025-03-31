package adapters

import (
	"ArquitecturaHexagonal_02/src/Loan/domain"
	"ArquitecturaHexagonal_02/src/helpers"
	"context"
	"encoding/json"
	"log"
	"time"
	"github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn *helpers.RabbitMQConn
}

// Constructor para Rabbit, manteniendo la conexión abierta
func NewRabbit() *Rabbit {
	conn, err := helpers.GetRabbitMQConn()
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
	}

	return &Rabbit{conn: conn}
}

// Método para enviar un mensaje a RabbitMQ
func (rabbit *Rabbit) PublishEvent(book domain.Book) error {
	err := rabbit.conn.Ch.ExchangeDeclare(
		"opt",   // Nombre del exchange
		"fanout", // Tipo de exchange (fanout enviará el mensaje a todas las colas vinculadas)
		true,     // Durable
		false,    // Auto-deleted
		false,    // Internal
		false,    // No-wait
		nil,      // Arguments
	)
	if err != nil {
		log.Printf("Error al declarar el exchange: %v", err)
		return err
	}

	// Declaración de la cola (esto solo debería hacerse una vez)
	_, err = rabbit.conn.Ch.QueueDeclare(
		"colaOpcional", // Nombre de la cola
		true,          // Durable
		false,         // Auto-delete
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Printf("Error al declarar la cola: %v", err)
		return err
	}

	// Vincular la cola al exchange
	err = rabbit.conn.Ch.QueueBind(
		"colaOpcional", // Nombre de la cola
		"",            // Routing key (vacío para fanout)
		"opt",        // Nombre del exchange
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Printf("Error al vincular la cola al exchange: %v", err)
		return err
	}

	// Convertir el libro a JSON
	body, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error al serializar el mensaje: %v", err)
		return err
	}

	// Contexto con timeout para la publicación
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publicar el mensaje al exchange (será enviado a todas las colas vinculadas)
	err = rabbit.conn.Ch.PublishWithContext(ctx,
		"opt", // Exchange
		"",     // Routing key (vacío para fanout)
		false,  // Mandatory
		false,  // Immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Printf("Error al publicar mensaje: %v", err)
		return err
	}

	log.Printf(" [x] Mensaje enviado al exchange, que será entregado a la cola: %s", body)
	return nil
}
