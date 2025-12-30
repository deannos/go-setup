package postgres

import (
	"context"
	"database/sql"
	"go-setup/internal/entity"
	"go-setup/pkg/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}
	if err != nil {
		return nil, errors.NewDatabaseError(err)
	}
	return user, nil
}

func (r *UserRepository) List(ctx context.Context) ([]*entity.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.NewDatabaseError(err)
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, errors.NewDatabaseError(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, name, email FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.ErrNotFound
	}
	if err != nil {
		return nil, errors.NewDatabaseError(err)
	}
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	if rows == 0 {
		return errors.ErrNotFound
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	if rows == 0 {
		return errors.ErrNotFound
	}
	return nil
}
