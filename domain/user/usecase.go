package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/luryon/go-ecommerce/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

func (u User) Create(m *model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.New()", err)
	}
	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte("{}")
	}

	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "u.storage.Create()", err)
	}

	m.Password = ""

	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "u.storage.GetByEmail()", err)
	}

	return user, nil
}

func (u User) GetByID(ID uuid.UUID) (model.User, error) {
	user, err := u.storage.GetByID(ID)
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	user.Password = ""

	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return model.Users{}, fmt.Errorf("%s %w", "u.storage.GetAll()", err)
	}

	return users, nil
}

func (u User) Login(email, password string) (model.User, error) {
	m, err := u.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "user.GetBuEmail()", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "bcrypt.CompareHashAndPassword()", err)
	}

	m.Password = ""

	return m, nil
}