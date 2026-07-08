package repository

import (
	"database/sql"
	"time"

	"github.com/aryeteinc/ferrotui/internal/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Create(p models.Product) error
	Update(p models.Product) error
	Delete(id int) error
}

type SQLiteProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &SQLiteProductRepository{db: db}
}

func (r *SQLiteProductRepository) GetAll() ([]models.Product, error) {
	q := "SELECT * FROM products"
	rows, err := r.db.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []models.Product{}

	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.StockQuantity,
			&product.CreatedAt,
			&product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *SQLiteProductRepository) GetByID(id int) (*models.Product, error) {
	q := "SELECT * FROM products WHERE id = ?"
	row := r.db.QueryRow(q, id)
	product := &models.Product{}
	if err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.StockQuantity,
		&product.CreatedAt,
		&product.UpdatedAt); err != nil {
		return nil, err
	}
	return product, nil
}

func (r *SQLiteProductRepository) Create(p models.Product) error {
	q := `INSERT INTO products
	(name, description, price, stock_quantity, created_at, updated_at)
	VALUES (?,?,?,?,?,?,?)`

	_, err := r.db.Exec(q, p.Name, p.Description, p.Price, p.StockQuantity, time.Now(), time.Now())
	return err
}

func (r *SQLiteProductRepository) Update(p models.Product) error {
	q := `UPDATE products SET
	name = ?,
	description =?,
	price = ?,
	stock_quantity =?,
	updated_at = ? WHERE id = ?`

	_, err := r.db.Exec(q, p.Name, p.Description, p.Price, p.StockQuantity, time.Now(), p.ID)
	return err
}

func (r *SQLiteProductRepository) Delete(id int) error {
	q := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(q, id)
	return err
}
