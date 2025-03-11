package store

import (
	"context"
	"log"
	"top1affiliate/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore interface {
	GetUserFromID(ctx context.Context, id string) (*models.User, error)
}

type userStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) UserStore {
	return &userStore{db: db}
}

func (s *userStore) GetUserFromID(ctx context.Context, id string) (*models.User, error) {

	var user models.User

	query := `SELECT id, affiliate_id, password, name FROM users WHERE affiliate_id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.AffiliateID, &user.Password, &user.Name); err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
