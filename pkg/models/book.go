package models

type Book struct {
	BookId   int64    `json:"id" gorm:"primaryKey"`
	Title    string   `json:"title"`
	Quantity int      `json:"quantity"`
	Genre    Genre    `json:"genre"`
	Author   string   `json:"author"`
	Review   []Review `json:"review"`
}

func newBook() Book {
	book := Book{}
	book.BookId = 1
	book.Title = "Hp"
	book.Quantity = 5
	book.Author = "JK Rowling"
	book.Genre = Genre{1, "fantasy"}
	book.Review = []Review{
		{
			1,
			"Good",
		},
	}
	return book
}
