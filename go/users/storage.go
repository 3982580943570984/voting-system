package users

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Create(user User) (int, error) {
	query := `
		INSERT INTO users (email, password, is_active)
		VALUES ($1, $2, $3)
		RETURNING user_id
	`

	var userID int

	error := s.db.QueryRow(query, user.Email, user.Password, user.IsActive).Scan(&userID)

	if error != nil {
		return 0, fmt.Errorf("error creating user: %w", error)
	}

	return userID, nil
}

func (s *Storage) GetAll() ([]User, error) {
	query := `
		SELECT user_id, email, password, is_active
		FROM users
	`

	rows, error := s.db.Query(query)

	if error != nil {
		return nil, fmt.Errorf("error retrieving users: %w", error)
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		user := User{}

		error = rows.Scan(&user.Id, &user.Email, &user.Password, &user.IsActive)

		if error != nil {
			return nil, fmt.Errorf("error scanning user: %w", error)
		}

		users = append(users, user)
	}

	error = rows.Err()

	if error != nil {
		return nil, fmt.Errorf("error with rows: %w", error)
	}

	return users, nil
}

func (s *Storage) GetByID(id int) (*User, error) {
	query := `
		SELECT user_id, email, password, is_active
		FROM users
		WHERE user_id = $1
	`
	user := &User{}

	error := s.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Password, &user.IsActive)

	if error != nil {
		if error == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error retrieving user: %w", error)
	}

	return user, nil
}

func (s *Storage) GetByEmail(email string) (*User, error) {
	query := `
		SELECT user_id, email, password, is_active
		FROM users
		WHERE email = $1
	`

	user := &User{}

	error := s.db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password, &user.IsActive)

	if error != nil {
		if error == sql.ErrNoRows {
			return nil, nil
		}
		return nil, error
	}

	return user, nil
}

func (s *Storage) Update(user User) error {
	query := `
		UPDATE users
		SET email = $1, password = $2, is_active = $3
		WHERE user_id = $4
	`
	_, err := s.db.Exec(query, user.Email, user.Password, user.IsActive, user.Id)

	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (s *Storage) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE user_id = $1
	`
	_, err := s.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}
