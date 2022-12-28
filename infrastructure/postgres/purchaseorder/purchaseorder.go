package purchaseorder

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luryon/go-ecommerce/infrastructure/postgres"
	"github.com/luryon/go-ecommerce/model"
)

const table = "purchase_orders"

var fields = []string{
	"id",
	"user_id",
	"products",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuilderSQLInsert(table, fields)

	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type PurchaseOrder struct {
	db *pgxpool.Pool
}

func New(dbpool *pgxpool.Pool) PurchaseOrder {
	return PurchaseOrder{dbpool}
}

func (p PurchaseOrder) Create(m *model.PurchaseOrder) error {
	_, err := p.db.Exec(context.Background(),
		psqlInsert,
		m.ID,
		m.UserID,
		m.Products,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p PurchaseOrder) GetByID(ID uuid.UUID) (model.PurchaseOrder, error) {
	query := psqlGetAll + " WHERE id = $1"

	row := p.db.QueryRow(context.Background(),
		query,
		ID,
	)

	return p.scanRow(row)
}

func (p PurchaseOrder) scanRow(s pgx.Row) (model.PurchaseOrder, error) {
	m := model.PurchaseOrder{}

	updateAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.UserID,
		&m.Products,
		&m.CreatedAt,
		&updateAtNull,
	)
	if err != nil {
		return model.PurchaseOrder{}, err
	}

	m.UpdatedAt = updateAtNull.Int64

	return m, nil
}
