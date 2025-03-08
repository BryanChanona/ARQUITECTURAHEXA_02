package infraestructure

import (
	"ArquitecturaHexagonal_02/src/Loan/domain"
	"database/sql"
	"log"
)

// Representa una conexión a la base de datos.
type MySQL struct {
	db *sql.DB
}

// Usamos esta función para crear una instancia de la estructura MySQL
func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(book domain.Book) error {

	query, err := mysql.db.Prepare("INSERT INTO booksAvailable (author, title, available) VALUES (?, ?,?)")

	if err != nil {
		return err
	}
	defer query.Exec()

	_, err = query.Exec(book.Author,book.Title,book.Available)

	if err != nil {
		log.Println("Error saving book:", err)
		return err
	}
	return nil
}
