package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)", id, name, description)
	if err != nil {
		return Category{}, err
	}
	// não precisa preencher todos os campos obrigatoriamente.
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) FindByCourseID(id string) (Category, error) {
	rows, err := c.db.Query("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = ?", id)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return Category{}, sql.ErrNoRows
	}

	var category Category
	if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return Category{}, err
	}
	return category, nil
}

func (c *Category) FindByID(id string) (Category, error) {
	row := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id)

	var category Category
	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		if err == sql.ErrNoRows {
			return Category{}, sql.ErrNoRows
		}
		return Category{}, err
	}

	return category, nil
}
