package user

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luryon/go-ecommerce/infrastructure/postgres"
	"github.com/luryon/go-ecommerce/model"
)

const table = "users"

var fields = []string{
	"id",
	"email",
	"password",
	"details",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuilderSQLInsert(table, fields)
	psqlGetAll = postgres.BuildSQLSelect(table, fields)
)

type User struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) User {
	return User{db}
}

func (u User) Create(m *model.User) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Email,
		m.Password,
		m.IsAdmin,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdateAt),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u User) GetByID(ID uuid.UUID) (model.User, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := u.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return u.scanRow(row, false)
}

func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlGetAll + "WHERE email = $1"
	row, err := u.db.Query(context.Background(), query, email)
	if err != nil {
		return model.User{}, nil
	}

	m, err := u.scanRow(row, true)
	if err != nil {
		return model.User{}, nil
	}

	return m, nil
}

func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(
		context.Background(),
		psqlGetAll)
	if err != nil {
		return nil, err
	}

	ms := model.Users{}
	for rows.Next() {
		m, err := u.scanRow(rows, false)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}

func (u User) scanRow(s pgx.Row, withPass bool) (model.User, error) {
	m := model.User{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.Email,
		&m.Password,
		&m.IsAdmin,
		&m.Details,
		&m.CreatedAt,
		&updatedAtNull)
	if err != nil {
		return model.User{}, err
	}
	m.UpdateAt = updatedAtNull.Int64

	if !withPass {
		m.Password = ""
	}

	return m, nil
}
