package infraestructure

import (
	"ArquitecturaHexagonal_02/src/Loan/domain"
	"database/sql"
	"fmt"
	"log"
)

// Representa una conexión a la base de datos.
type MySQL struct {
	db *sql.DB
	lastCount int
}

// Usamos esta función para crear una instancia de la estructura MySQL
func NewMySQL(db *sql.DB) *MySQL {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM booksavailable").Scan(&count)
	if err != nil {
		log.Fatal("Error obteniendo el conteo inicial: ", err)
	}

	// Inicializamos el repositorio con el conteo inicial
	return &MySQL{
		db:        db,
		lastCount: count, // Establecemos el conteo inicial
	}
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

func (mysql *MySQL) GetBooks()([]domain.Book,error)  {
	query, err := mysql.db.Query("SELECT * FROM booksavailable")

	if err != nil {
		fmt.Println("Error getting books:", err)
	}

	var books []domain.Book
	defer query.Close()
			var book domain.Book
		for query.Next() {
			err := query.Scan(&book.Id,&book.Author, &book.Title, &book.Available)
			if err != nil {
				fmt.Println("Error getting books:", err)
			}
			fmt.Println(book)
			books = append(books,book)
		}
		return books, nil

}

func (sql *MySQL) GetNewBookIdAdded() (bool, []domain.Book, error) {
	var count int
	err := sql.db.QueryRow("SELECT COUNT(*) FROM booksavailable").Scan(&count)
	if err != nil {
		return false, nil, fmt.Errorf("error obteniendo el conteo de libros: %v", err)
	}

	if count > sql.lastCount {
		// Actualizamos lastCount para mantener el conteo más reciente
		sql.lastCount = count

		// Obtener la lista actualizada de libros
		rows, err := sql.db.Query("SELECT id, title, author, available FROM booksavailable")
		if err != nil {
			return false, nil, fmt.Errorf("error obteniendo la lista de libros: %v", err)
		}
		defer rows.Close()

		var books []domain.Book
		for rows.Next() {
			var book domain.Book
			if err := rows.Scan(&book.Id,&book.Title, &book.Author,&book.Available); err != nil {
				return false, nil, fmt.Errorf("error escaneando los libros: %v", err)
			}
			books = append(books, book)
		}

		return true, books, nil
	}

	return false, nil, nil
}


