package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  int
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description string, categoryID int) (Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryID)

	if err != nil {
		return Course{}, err
	}

	return Course{db: c.db, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		var id, name, description string
		var categoryID int

		err := rows.Scan(&id, &name, &description, &categoryID)

		if err != nil {
			return nil, err
		}

		courses = append(courses, Course{db: c.db, ID: id, Name: name, Description: description, CategoryID: categoryID})
	}

	return courses, nil
}
