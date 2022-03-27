package book

import (
	"Quiz-3/config"
	"Quiz-3/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table          = "book"
	layoutDateTime = "2022-03-27 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		// var createdAt, updatedAt string

		if err = rowQuery.Scan(&book.Id,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Created_at,
			&book.Update_at,
			&book.Category_id,
		); err != nil {
			return nil, err
		}

		// //  Change format string to datetime for created_at and updated_at
		// book.Created_at, err = time.Parse(layoutDateTime, createdAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// book.Update_at, err = time.Parse(layoutDateTime, updatedAt)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		books = append(books, book)
	}
	return books, nil
}

func GetTitle(ctx context.Context, title string)([]models.Book, error){
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Where title = %s", table, title)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		// var createdAt, updatedAt string

		if err = rowQuery.Scan(&book.Id,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Created_at,
			&book.Update_at,
			&book.Category_id,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetMinYear(ctx context.Context, year int)([]models.Book, error){
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Where year >= %d", table, year)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		// var createdAt, updatedAt string

		if err = rowQuery.Scan(&book.Id,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Created_at,
			&book.Update_at,
			&book.Category_id,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBookFromCategory(ctx context.Context, id string)([]models.Book, error){
	var books []models.Book
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Where category_id = %s Order By created_at DESC", table, id)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		// var createdAt, updatedAt string

		if err = rowQuery.Scan(&book.Id,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Created_at,
			&book.Update_at,
			&book.Category_id,
		); err != nil {
			return nil, err
		}

		// //  Change format string to datetime for created_at and updated_at
		// book.Created_at, err = time.Parse(layoutDateTime, createdAt)

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// book.Update_at, err = time.Parse(layoutDateTime, updatedAt)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		books = append(books, book)
	}
	return books, nil
}

// Insert Book
func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) values('%v', '%v', '%v', %v, '%v', %v,'%v', NOW(), NOW(), %v)",
	table, 
	book.Title,
	book.Description,
	book.Image_url,
	book.Release_year,
	book.Price,
	book.Total_page,
	book.Thickness,
	book.Category_id,
)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Book
func Update(ctx context.Context, book models.Book, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set tite ='%s', description = '%s', image_url = '%s', release_year = %d, price = '%s', total_page = %d, category_id = %d updated_at = NOW() where id = %s",
		table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_year,
		book.Price,
		book.Total_page,
		book.Category_id,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Category
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}