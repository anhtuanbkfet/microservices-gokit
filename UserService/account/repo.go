package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var (
	RepoErr = errors.New("Unable to handle Repo Request")
)

type userRepository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &userRepository{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *userRepository) CreateUser(ctx context.Context, user User) error {
	sql := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)
			ON CONFLICT (id) WHERE (email = $2) DO NOTHING;`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.Id, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return "", RepoErr
	}

	return email, nil
}

/*
@ return account ID if email and password is correct; or NIL and error code if failed
*/
func (repo *userRepository) ValidateUser(ctx context.Context, email, password string) (string, error) {
	var id string
	err := repo.db.QueryRow("SELECT id FROM users WHERE email=$1 AND password=$2", email, password).Scan(&id)
	if err != nil {
		return "", ErrInvalidUser
	}

	return id, nil
}
