package store

import (
	"database/sql"
	"go-comments-api/internal/model"
	"time"
)

type Store interface {
	GetAll() ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(id int, user *model.User) (*model.User, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s store) GetAll() ([]*model.User, error) {
	q := "SELECT id, name, lastName, userName, email, age, status, created_At, updated_At FROM users WHERE status = 1"

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		u := model.User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.LastName, &u.Username, &u.Email, &u.Age, &u.Status, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (s *store) GetByID(id int) (*model.User, error) {
	q := `SELECT id, name, lastName, userName, email, age, status, created_At, updated_At FROM users WHERE id = ? AND status = 1`

	u := model.User{}

	err := s.db.QueryRow(q, id).Scan(&u.ID, &u.Name, &u.LastName, &u.Username, &u.Email, &u.Age, &u.Status, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *store) Create(user *model.User) (*model.User, error) {
	q := "INSERT INTO users (name, lastName, userName, password, email, age, status, created_At) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	user.Status = 1
	user.CreatedAt = time.Now()

	resp, err := s.db.Exec(q, user.Name, user.LastName, user.Username, user.Password, user.Email, user.Age, user.Status, user.CreatedAt)
	if err != nil {
		return nil, err
	}
	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)

	return user, nil
}

func (s *store) Update(id int, user *model.User) (*model.User, error) {
	q := `UPDATE users SET name = ?, lastName = ?, age = ?, updated_At = ? WHERE id = ?`

	user.UpdatedAt = time.Now()

	_, err := s.db.Exec(q, user.Name, user.LastName, user.Age, user.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (s *store) Delete(id int) error {
	q := `UPDATE users SET status = 0 WHERE id = ?`

	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}
