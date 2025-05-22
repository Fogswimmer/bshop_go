package seeder

import (
	"api/train/models/entities"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jaswdr/faker"
)

func Seed(db *sql.DB) error {
	log.Println("Seeding database...")
	err := truncateTables(db)
	if err != nil {
		return fmt.Errorf("TruncateTables error: %v", err)
	}
	authors, err := seedAuthors(db)
	if err != nil {
		log.Println("Authors seeding error: ", err)
		return fmt.Errorf("AuthorSeeding error: %v", err)
	}
	err = seedBooks(db, authors)
	if err != nil {
		log.Println("Books seeding error: ", err)
		return fmt.Errorf("BookSeeding error: %v", err)
	}
	log.Println("Seeding database completed successfully!")
	return nil
}

func seedBooks(db *sql.DB, authors []entities.Author) (err error) {
	fake := faker.New()
	log.Println("Books seeding started...")
	for i := 0; i < 10; i++ {
		a := authors[fake.IntBetween(0, len(authors)-1)]
		snt := fake.Lorem().Sentence(2)
		title := strings.ToUpper(snt[:1]) + strings.Trim(snt[1:], ".")
		relYear := fake.IntBetween(1900, 2021)
		p := fake.Lorem().Paragraph(5)
		sum := strings.ToUpper(p[:1]) + p[1:]
		pr := fake.Float64(0, 99, 2)
		res, err := db.Exec(
			"INSERT INTO book (title, release_year, summary, price, author_id) VALUES ($1, $2, $3, $4, $5)",
			title, relYear, sum, pr, a.ID)
		if err != nil {
			log.Println("Books seeding error: ", err)
			return fmt.Errorf("SeedBooks error: %v", err)
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Println("RowsAffected error: ", err)
			return fmt.Errorf("RowsAffected error: %v", err)
		}
		if rowsAffected == 0 {
			log.Println("no books were inserted: ", err)
			return fmt.Errorf("no books were inserted")
		}
		log.Printf("Book %d seeded: %s %d %s %f %s", i+1, title, relYear, sum, pr, a.Lastname)
	}
	log.Println("Books seeding finished")
	return nil
}

func seedAuthors(db *sql.DB) (authors []entities.Author, err error) {
	fake := faker.New()
	authors = make([]entities.Author, 10)
	log.Println("Authors seeding started...")
	for i := range authors {
		fname := fake.Person().FirstName()
		lname := fake.Person().LastName()
		bday := fake.Time().TimeBetween(time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC), time.Now())

		var id int
		err := db.QueryRow("INSERT INTO author (firstname, lastname, birthday) VALUES ($1, $2, $3) RETURNING id", fname, lname, bday).Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("SeedAuthors error: %v", err)
		}

		bdayStr := bday.Format("2006-01-02")
		authors[i] = entities.Author{
			ID:        id,
			Firstname: fname,
			Lastname:  lname,
			Birthday:  bdayStr,
		}
		log.Printf("Author %d seeded: %s %s %s (ID: %d)", i+1, fname, lname, bdayStr, id)
	}
	log.Println("Authors seeding finished")
	return authors, nil
}

func truncateTables(db *sql.DB) error {
	tables := []string{"book", "author"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			return fmt.Errorf("TruncateTables error: %v", err)
		}
		log.Printf("Truncated table %s", table)
	}
	return nil
}
