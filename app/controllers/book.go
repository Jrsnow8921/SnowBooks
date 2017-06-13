package controllers

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	models "SnowGo/app/models"
)

func getBook(bookID int) (models.Book, error) {
	//Retrieve
	res := models.Book{}

	var id int
	var name string
	var author string
	var pages int
	var publicationDate pq.NullTime

	err := db.QueryRow(`SELECT id, name, author, pages, publication_date FROM books where id = $1`, bookID).Scan(&id, &name, &author, &pages, &publicationDate)
	if err == nil {
		res = models.Book{ID: id, Name: name, Author: author, Pages: pages, PublicationDate: publicationDate.Time}
	}

	return res, err
}

func getBookLastID() int { 

    var id int

	err := db.QueryRow("select ifnull(max(id), 0) as id from books").Scan(&id)
    if err != nil {
      panic(err)
    }
    fmt.Println("New record ID is:", id + 1)
    return id + 1
}

func allBooks() ([]models.Book, error) {
	//Retrieve
	books := []models.Book{}

	rows, err := db.Query(`SELECT id, name, author, pages, publication_date FROM books order by id`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var author string
			var pages int
			var publicationDate pq.NullTime

			err = rows.Scan(&id, &name, &author, &pages, &publicationDate)
			if err == nil {
				currentBook := models.Book{ID: id, Name: name, Author: author, Pages: pages}
				if publicationDate.Valid {
					currentBook.PublicationDate = publicationDate.Time
				}

				books = append(books, currentBook)
			} else {
				return books, err
			}
		}
	} else {
		return books, err
	}

	return books, err
}

func insertBook(name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
        
    lastid := getBookLastID() 
                
        
	res, err := db.Exec(`INSERT INTO books(id, name, author, pages, publication_date) VALUES($1, $2, $3, $4, $5)`, lastid, name, author, pages, publicationDate)

	if err != nil {
		return 0, err
	}

	rowsCreated, err := res.RowsAffected()

		if err != nil {
		return 0, err
	}

	return int(rowsCreated), err

}

func updateBook(id int, name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE books set name=$1, author=$2, pages=$3, publication_date=$4 where id=$5`, name, author, pages, publicationDate, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeBook(bookID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from books where id = $1`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
