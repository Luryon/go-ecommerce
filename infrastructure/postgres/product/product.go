package product

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luryon/go-ecommerce/infrastructure/postgres"
	"github.com/luryon/go-ecommerce/model"
)

const table = "products"

var fields = []string{
	"id",
	"product_name",
	"price",
	"images",
	"description",
	"features",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuilderSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdateByID(table, fields)
	PSQLDelete = postgres.BuildSQLDelete(table)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type Product struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Product {
	return Product{db}
}

func (p Product) Create(m *model.Product) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.ProductName,
		m.Price,
		m.Images,
		m.Description,
		m.Features,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p Product) Update(m *model.Product) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		m.ProductName,
		m.Price,
		m.Images,
		m.Description,
		m.Features,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p Product) Delete(ID uuid.UUID) error {
	_, err := p.db.Exec(context.Background(),
		PSQLDelete,
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	query := psqlGetAll + "WHERE id = $1"
	row := p.db.QueryRow(
		context.Background(),
		query,
		ID,
	)
	return p.scanRow(row)
}

func (p Product) GetAll() (model.Products, error) {
	rows, err := p.db.Query(context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Products
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (p Product) scanRow(s pgx.Row) (model.Product, error) {
	m := model.Product{}

	updateAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.ProductName,
		&m.Price,
		&m.Images,
		&m.Description,
		&m.Features,
		&m.CreatedAt,
		&updateAtNull,
	)
	if err != nil {
		return model.Product{}, err
	}

	m.UpdatedAt = updateAtNull.Int64

	return m, nil
}
