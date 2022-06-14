package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"mymodule/pkg/models"
)

// func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
// 	var books []models.Book

// 	if result := h.DB.Find(&books); result.Error != nil {
// 		fmt.Println(result.Error)
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(books)
// }

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	books, err := getAllBooks()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(books)
}

// get one user from the DB by its userid
func getAllBooks() ([]models.Book, error) {
	// create an empty Book of type models.Book
	var book models.Book

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()
	paramCount := 0

	var books []models.Book

	// create the select sql query
	sqlStatement := `SELECT id, title, quantity, author, coalesce(genre_id, 0), coalesce(review_id, 0) FROM books`

	if &book.Author != nil || &book.Genre != nil {
		sqlStatement = sqlStatement + " where "

		if &book.Author != nil {
			sqlStatement = sqlStatement + " author = $1"
			paramCount++
		}

		if paramCount > 0 {
			sqlStatement = sqlStatement + " and "
		}

		if &book.Genre != nil {

			//	validateGenre := `SELECT book_id FROM genres where genre = $1`

			// execute the sql statement
			//	rows, _ := db.Query(validateGenre, genre.Genre)
			sqlStatement = sqlStatement + " "
		}
	}

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var book models.Book

		// unmarshal the row object to user
		err = rows.Scan(&book.Id, &book.Title, &book.Quantity, &book.Author, &book.GenreId, &book.ReviewId)

		sqlGenreStatement := `SELECT id, genre from genres where book_id = $1`
		genreRows, err := db.Query(sqlGenreStatement, &book.Id)
		var genre models.Genre
		for genreRows.Next() {

			err = genreRows.Scan(&genre.Id, &genre.Genre)

		}
		book.Genre = genre
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		books = append(books, book)

	}

	// return empty user on error
	return books, err
}
