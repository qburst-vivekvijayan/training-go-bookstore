package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"mymodule/pkg/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qburst123"
	dbname   = "book_store"
)

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	//db, err := sql.Open("postgres", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// func (h handler) Store(w http.ResponseWriter, r *http.Request) {
// 	// Read to request body
// 	defer r.Body.Close()
// 	body, err := ioutil.ReadAll(r.Body)

// 	ret := "Created"

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var book models.Book
// 	json.Unmarshal(body, &book)

// 	// Append to the Books table
// 	if result := h.DB.Create(&book); result.Error != nil {
// 		fmt.Println(result.Error)
// 		ret = string(result.Error.Error())
// 	}

// 	// Send a 201 created response
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(ret)
// }

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty Book of type models.Book
	var book models.Book

	// decode the json request to Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	var genreId int64 = 0

	if book.Genre != (models.Genre{}) {
		genre := book.Genre
		genreId = insertGenre(genre)
	}

	// call insert Book function and pass the Book
	insertID := insertBook(book, genreId)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Book added successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func insertBook(book models.Book, genreId int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning Bookid will return the id of the inserted Book
	sqlStatement := `INSERT INTO Books (title, quantity, author) VALUES ($1, $2, $3) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, book.Title, book.Quantity, book.Author).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute insertBook query book. %v", err)
	}

	if book.Genre != (models.Genre{}) {
		genre := book.Genre
		updateGenre(genre, id)
	}

	fmt.Printf("Inserted a single book record %v", id)

	// return the inserted id
	return id
}

func insertGenre(genre models.Genre) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning Bookid will return the id of the inserted Book

	validateGenre := `SELECT * FROM genres where genre = $1`

	// execute the sql statement
	rows, _ := db.Query(validateGenre, genre.Genre)

	counter := 0
	for rows.Next() {
		//get row count
		counter++
	}
	log.Println("Counter :", counter)
	var id int64
	if counter == 0 {

		sqlStatement := `INSERT INTO genres (genre) VALUES ($1) RETURNING id`

		// the inserted id will store in this id

		// execute the sql statement
		// Scan function will save the insert id in the id
		err := db.QueryRow(sqlStatement, genre.Genre).Scan(&id)

		if err != nil {
			log.Fatalf("Unable to execute insertGenre the query genre. %v", err)
		}

		fmt.Printf("Inserted a single genre %v", id)
	}
	// return the inserted id
	return id
}

func updateGenre(genre models.Genre, book_id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning Bookid will return the id of the inserted Book
	sqlStatement := `update genres set book_id = $1 where genre = $2`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	_, err := db.Exec(sqlStatement, book_id, genre.Genre)

	if err != nil {
		log.Fatalf("Unable to execute the query genre. %v", err)
	}

	fmt.Printf("updated genre %v", id)

	// return the inserted id
	return id
}
